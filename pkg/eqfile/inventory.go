// Copyright 2020 Nuttann. All rights reserved.
// The use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package eqfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// InvItem is the data found in each line of an Inventory file.
type InvItem struct {
	Loc   string // Inventory/bank/bag slot.
	Name  string // Item name or "Empty"
	ID    int    // Item ID or 0
	Count int    // Number of that item
	Slots int    // Number of slots if a bag, 6 if another item, or 0 if empty
}

// Expected inventory header line
const inventoryHeader = "Location\tName\tID\tCount\tSlots"

// ReadInventory will return an array of Items from the inventory file.
//
// The format appears to be five columns (Location, Name, ID, Count, and Slots)
// separated by a tab character. The first line is a header line.
//
// This will do some simple checking of the file format and if any issues are
// identified, it will return an error instead of any data.
//
// Error reasons:
//   - File cannot be opened for reading.
//   - Header line does not match expected header.
//   - Column 3, 4, or 5 are not integers.
func ReadInventory(fname string) ([]InvItem, error) {
	var items []InvItem
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Check header line
	_ = scanner.Scan()
	head := scanner.Text()
	if head != inventoryHeader {
		// Unexpected header line. Either not an inventory file
		// or format may have changed.
		return nil, fmt.Errorf("%s: missing expected header of \"%s\"",
			fname, inventoryHeader)
	}
	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		parts := strings.Split(line, "\t")
		if len(parts) != 5 {
			return nil, fmt.Errorf("%s at line %d has %d columns instead of 5",
				fname, lineNo, len(parts))
		}
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 3 is not an integer",
				fname, lineNo)
		}
		count, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 4 is not an integer",
				fname, lineNo)
		}
		slots, err := strconv.Atoi(parts[4])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 5 is not an integer",
				fname, lineNo)
		}
		items = append(items, InvItem{
			Loc:   parts[0],
			Name:  parts[1],
			ID:    id,
			Count: count,
			Slots: slots,
		})
	}
	return items, err
}
