// Copyright 2020 Nuttann. All rights reserved.
// The use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/nuttann/equtils/pkg/eqdb"
	"github.com/nuttann/equtils/pkg/eqfile"
	"gopkg.in/yaml.v2"
)

// 'config' holds locations of necessary data files.
type config struct {
	ItemDBLoc   string   // DB info (currently file name)
	RealEstate  []string // All Real Estate Files to be read
	Inventories []string // All Inventory Files to be read
	LootFilters []string // All Loot Filter Files to be read
}

func readConfig(confFile string) (confData config) {
	// Read configuration data.
	buf, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatalf("error: Configuration file - %v", err)
	}
	err = yaml.Unmarshal(buf, &confData)
	if err != nil {
		log.Fatalf("error: Configuration file - %v", err)
	}
	return
}

func main() {
	confPtr := flag.String("conf", "", "Configuration file. (Required)")
	flag.Parse()

	if *confPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// Get file paths to necessary data files.
	conf := readConfig(*confPtr)

	itemDB := eqdb.OpenItemDB(conf.ItemDBLoc)
	defer itemDB.Close()

	for _, re := range conf.RealEstate {
		reData, err := eqfile.ReadRE(re)
		if err != nil {
			log.Fatalf("error: Reading house file - %v", err)
		}
		for _, entry := range reData {
			itemDB.SetName(entry.ID, entry.ItemName)
		}
	}
	for _, i := range conf.Inventories {
		// Process Inventory File.
		invData, err := eqfile.ReadInventory(i)
		if err != nil {
			log.Fatalf("error: Reading inventory file - %v", err)
		}
		for _, entry := range invData {
			itemDB.SetName(entry.ID, entry.Name)
		}
	}
	for _, lf := range conf.LootFilters {
		// Process LootFilters.
		lfData, err := eqfile.ReadLF(lf)
		if err != nil {
			log.Fatalf("error: Reading lootfilter file - %v", err)
		}
		for _, entry := range lfData {
			// Loot filter names may be old. Only add new names.
			if itemDB.Name(entry.ID) == "" {
				itemDB.SetName(entry.ID, entry.Name) // Use this if name is not known.
			}
			itemDB.SetIconID(entry.ID, entry.IconID)
		}
	}
}
