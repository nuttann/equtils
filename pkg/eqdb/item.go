// Copyright 2020 Nuttann. All rights reserved.
// The use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package eqdb

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Item is the EQ item's data that is known and useful.
type Item struct {
	Name   string
	IconID int
}

// Items is the known items and associated useful attributes.
type Items struct {
	DB      map[int]Item // Currently known items
	Fname   string       // File to hold DB
	Changed bool         // Set to true if DB is altered and should be saved.
}

// OpenItemDB will return a DB read in from a YAML file.
func OpenItemDB(dbFile string) (i Items) {
	i.Fname = dbFile
	dat, err := ioutil.ReadFile(dbFile)
	if err != nil {
		fmt.Print("Cannot read DB file. Starting with empty item DB.")
		i.DB = make(map[int]Item)
		return
	}
	_ = yaml.Unmarshal(dat, &i.DB)
	return
}

// Close will save the item DB to a YAML file if the DB changed.
func (i *Items) Close() {
	if i.Changed {
		fmt.Println("New data - Updating DB file.")
		fmt.Printf("    Orig DB file name - %s\n", i.Fname)
		fmt.Printf("    DB directory %s\n", filepath.Dir(i.Fname))
		f, err := ioutil.TempFile(filepath.Dir(i.Fname), "tempdb")
		if err != nil {
			fmt.Printf("Could not get temp name - %v\n", err)
			os.Exit(0)
		}
		fmt.Printf("    Temporary DB file - %s\n", f.Name())
		defer os.Remove(f.Name())
		defer f.Close()
		if err != nil {
			fmt.Println("Error - Could not open temp DB file:", f.Name())
			return
		}
		dat, err := yaml.Marshal(&i.DB)
		if err != nil {
			fmt.Println("Error - Could not marshal item DB items.")
			return
		}
		_, err = f.Write(dat)
		if err != nil {
			fmt.Println("Error - could not write to temp DB file.")
			return
		}
		f.Close() // Must be closed before Rename().
		err = os.Rename(f.Name(), i.Fname)
		if err != nil {
			fmt.Printf("Error - Could not replace original DB file. %v", err)
		}
		fmt.Println("    Overwrote original item DB file.")
		return
	}
}

// Name will get the id's item Name.
func (i *Items) Name(id int) string {
	return i.DB[id].Name
}

// IconID will get the id's item IconID.
func (i *Items) IconID(id int) int {
	return i.DB[id].IconID
}

// GetItem will get the id's whole Item. It will return an empty item if it
// does not exist.
func (i *Items) GetItem(id int) Item {
	return i.DB[id]
}

// SetName will set the id's item Name.
func (i *Items) SetName(id int, name string) {
	item := i.DB[id]
	if item.Name != name {
		// Only set if different.
		item.Name = name
		i.DB[id] = item
		i.Changed = true
	}
}

// SetIconID will set the Icon ID for the item ID.
func (i *Items) SetIconID(id int, iconID int) {
	item := i.DB[id]
	if item.IconID != iconID {
		// Only set if different.
		item.IconID = iconID
		i.DB[id] = item
		i.Changed = true
	}
}

// SetItem will set the whole Item for the item ID.
func (i *Items) SetItem(id int, item Item) {
	if i.DB[id] != item {
		// Only set if different.
		i.DB[id] = item
		i.Changed = true
	}
}
