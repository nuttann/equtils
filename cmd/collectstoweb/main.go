// Copyright 2020 Nuttann. All rights reserved.
// The use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/brianholland99/intlist"
	"github.com/nuttann/equtils/pkg/eqdb"
	"github.com/nuttann/equtils/pkg/eqfile"
	"gopkg.in/yaml.v2"
)

// 'config' holds configuration read from the configuration file. This matches
// the YAML file format so that the YAML package can populate this. This
// includes the nested House definition that configures which quests'
// collection items are stored in each house.
type config struct {
	QuestsFile string  // File with exp - zone - quest - item ID mappings
	ItemDBLoc  string  // DB location info (currently file name)
	HTMLOut    string  // Where to write output
	HTMLTitle  string  // Single line title for header and <h1> tag
	HTMLIntro  string  // HTML to include between the <body> tag and the quest info
	Houses     []House // House collection configuration
}

// 'intState' holds internal state needed throughout the program. This includes
// I/O and database handles.
type intState struct {
	ItemDB    *eqdb.Items   // Handle to open DB (set when opening ItemDBLoc)
	QuestData *[]QuestExp   // Handle to quest data
	Out       *os.File      // HTML output to close
	Buf       *bufio.Writer // HTML output writer
}

// itemInfo collects counts for an item in a house while examining a realestate
// dump.
type itemInfo struct {
	Name   string
	Count  int // Total count of that item
	Stacks int // Number of slots with this item
}

// itemMap holds the counts for all items in a house.
type itemMap map[int]itemInfo

// There are similarities between generic quest data and specifiic house
// configuration fields. The names are prefixed with Quest and House
// respectively. These two sets of hierarchical structures will be merged and
// processed to produce the results for the output. These are filled directly
// from the YAML data files.

// QuestInfo holds info for a specific collection quest.
type QuestInfo struct {
	Name string // Quest name (E.g., "Fear in Pieces")
	Ids  string // EQ Item IDs contained in quest (E.g., "500-506, 519")
	Note string // Notes about this collection set (E.g, "Drop from undead")
}

// QuestZone holds collection quests for a zone.
type QuestZone struct {
	Name   string      // Zone name (E.g., "Shard's Landing")
	Quests []QuestInfo // All collection quests in this zone.
}

// QuestExp holds the zones for that expansion. This is the top-level quest
// structure. Expansions -> Zones -> Quests.
type QuestExp struct {
	Name  string      // Expansion name (E.g., "Rain of Fear")
	Zones []QuestZone // All zones in expansion with collection quests.
}

// HouseExp describes which expansions and zones within that expansion are
// contained in a house. The Name string here and the Name strings in the Zones
// array must match exactly with the Expansion.name and Zone.name fields in the
// quest set of structures.
type HouseExp struct {
	Name  string   // Expansion name
	Zones []string // Zones in expansion for this house. Empty means all zones.
}

// House holds the metadata and the configured expansions/zones for the house.
type House struct {
	Fname      string // Path/File name (E.g., "Nuttann_cazic-RealEstate.txt")
	Address    string // Property string (From the "/output realestate" file.)
	Expansions []HouseExp
}

// Function setup will read the configuration file and perform setup. This
// includes validating the data and reading data sources.
//
// Note: teardown() needs to be called to save any accumulated info such as new
// item data and to flush and close output.
func setup(confFile string) (intState, config) {
	// Read configuration data.
	var conf config
	var state intState
	buf, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatalf("error: Configuration file - %v", err)
	}
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Fatalf("error: Configuration file - %v", err)
	}
	questData := readQuests(conf.QuestsFile)
	state.QuestData = &questData
	itemDB := eqdb.OpenItemDB(conf.ItemDBLoc)
	state.ItemDB = &itemDB
	state.Out, err = os.Create(conf.HTMLOut)
	if err != nil {
		log.Fatalf("error: Opening output file - %v", err)
	}
	state.Buf = bufio.NewWriter(state.Out)
	return state, conf
}

// Function teardown will cleanup data files.
func teardown(state intState) {
	state.ItemDB.Close() // Updates if anything was changed.
	// Finish HTML output.
	state.Buf.Flush()
	state.Out.Close()
}

// readQuests reads a questfile and populates []QuestExp.
func readQuests(questsFile string) (questData []QuestExp) {
	buf, err := ioutil.ReadFile(questsFile)
	if err != nil {
		log.Fatalf("error: Quest data file - %v", err)
	}
	err = yaml.Unmarshal(buf, &questData)
	if err != nil {
		log.Fatalf("error: Quest data file - %v", err)
	}
	return
}

// getStoredItemData will return items from Real Estate dump for the given
// property. This only pulls out 'Stored' items since collection items
// are not placeable.
func getStoredItemData(house House, itemDB *eqdb.Items) itemMap {
	houseData, err := eqfile.ReadRE(house.Fname)
	if err != nil {
		log.Fatalf("error: Reading house file - %v", err)
	}
	items := make(itemMap)
	for _, entry := range houseData {
		id := entry.ID
		itemDB.SetName(id, entry.ItemName)
		if entry.RELoc != "Plot" || entry.Status != "Stored" {
			continue
		}
		if entry.REName != house.Address {
			continue
		}
		count := entry.Count
		if _, ok := items[id]; ok {
			items[id] = itemInfo{
				Count:  count + items[id].Count,
				Name:   items[id].Name,
				Stacks: items[id].Stacks + 1,
			}
		} else {
			items[id] = itemInfo{
				Count:  count,
				Name:   entry.ItemName,
				Stacks: 1,
			}
		}
	}
	return items
}

// writeZone will output all collection quests / items for a zone with a quest
// being expandable.
// Returns:
//    haveItems - # of unique quest items for this zone present in house.
//    totalItems - total # of unique quest items for this zone.
func writeZone(state intState, house House, zone QuestZone, items itemMap) (haveItems int, totalItems int) {
	state.Buf.WriteString("<li>" + zone.Name + "<ul>")
	for _, quest := range zone.Quests {
		ids, err := intlist.Parse(quest.Ids)
		if err != nil {
			log.Fatalf("error: Quest file bad int range - %v", err)
		}
		have := 0
		// Collect quest HTML in Buffer so that it can be added after summary
		// that includes counts. (Browsers vary display order if summary is at
		// end of details.)
		var buf bytes.Buffer
		for _, id := range ids {
			name := state.ItemDB.Name(id)
			if name == "" {
				name = "???" // Use this if name is not known.
			}
			if items[id].Stacks > 1 {
				fmt.Println("Multiple stacks -", items[id].Name, "-", house.Address)
			}
			count := items[id].Count
			if count > 0 {
				delete(items, id)
				have++ // Keep count of unique items for this quest in the house.
			}
			str := fmt.Sprint("<li>", name, " (", count, ")</li>\n")
			buf.WriteString(str)
		}
		state.Buf.WriteString("<details><summary>" + quest.Name + " (" + strconv.Itoa(have) + "/" + strconv.Itoa(len(ids)) + ") ")
		if quest.Note != "" {
			state.Buf.WriteString(" - " + quest.Note)
		}
		state.Buf.WriteString("\n</summary>\n<ul>\n" + buf.String() + "</ul>\n</details>\n")
		// Add counts to totals for zone.
		totalItems += len(ids)
		haveItems += have
	}
	state.Buf.WriteString("</ul></li>\n")
	return
}

// writeHouseHTML will output the HTML for the quests in the passed house.
func writeHouseHTML(state intState, house House) {
	state.Buf.WriteString("<h2>" + house.Address + "</h2><ul>\n")
	totalItems := 0
	haveItems := 0

	items := getStoredItemData(house, state.ItemDB) // 'items' is filtered to only have "Stored" items.
	for _, houseExp := range house.Expansions {
		for _, questExp := range *state.QuestData {
			if houseExp.Name == questExp.Name { // Found the quest data for this expansion.
				state.Buf.WriteString("<li>" + questExp.Name + "<ul>\n")
				if houseExp.Zones == nil {
					// Do all zones in this expansion.
					for _, questZone := range questExp.Zones {
						have, total := writeZone(state, house, questZone, items)
						totalItems += total
						haveItems += have
					}
				} else {
					// Do specific zones listed in order listed.
					for _, loczone := range houseExp.Zones {
						// Find the zone in quests.
						for _, questZone := range questExp.Zones {
							if loczone == questZone.Name {
								have, total := writeZone(state, house, questZone, items)
								totalItems += total
								haveItems += have
								break // Handled matching zone.
							}
						}
					}
				}
				state.Buf.WriteString("</ul></li>\n")
				break // Handled matching expansion.
			}
		}
	}
	state.Buf.WriteString("</ul><p>Summary - Items = " + strconv.Itoa(haveItems) + " / " + strconv.Itoa(totalItems) + "</p>\n")
	if len(items) > 0 {
		fmt.Println("====== Extra items in -", house.Address)
		for _, item := range items {
			fmt.Println(item.Name, " (", item.Count, ")")
		}
		fmt.Println()
	}
	return
}

// headTemplateDef is the html/template definition for the beginning of the
// HTML output. HTMLTitle is treated as text and escaped for HTML. HTMLIntro is
// treated as raw HTML, so the user can embed links and such.
const headTemplateDef = `<!DOCTYPE html>
<html>
    <head><title>{{.HTMLTitle}}</title></head>
    <body>
	    <h1>{{.HTMLTitle}}</h1>
		{{unescape .HTMLIntro}}
		<p>Click quest to expand items.
			The counts for items are in parentheses.</p>`

// writeHTML controls the overall HTML output.
func writeHTML(state intState, conf config) {
	// Output beginning HTML. HTML escape all except the 'HTMLIntro" config.
	mainTemplate, err := template.New("page").Funcs(template.FuncMap{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).Parse(headTemplateDef)
	if err != nil {
		log.Fatal(err)
	}
	err = mainTemplate.Execute(state.Buf, conf)
	if err != nil {
		log.Fatal(err)
	}
	// Process and output house data.
	for _, house := range conf.Houses {
		writeHouseHTML(state, house)
	}
	// Output end of HTML doc.
	state.Buf.WriteString("</body></html>\n")
	return
}

func main() {
	// Get configuration file name.
	confFile := flag.String("conf", "", "Configuration file. (Required)")
	flag.Parse()
	if *confFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Use configuration to setup DBs and outputs.
	state, conf := setup(*confFile)
	defer teardown(state) // Save any output and close files at end.

	writeHTML(state, conf)
}
