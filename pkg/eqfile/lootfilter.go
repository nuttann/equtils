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

// LFItem is the data found in each line of a Loot Filter file.
type LFItem struct {
	ID     int    // Item ID (uniquely identifies an item)
	IconID int    // The ID of the corresponding icon to display in advloot.
	Name   string // Name of item at the time its setting was made.
}

// Expected loot-filter header line
const lfHeader = "#ITEM_ID^ICON_ID^ITEM_NAME"

// ReadLF will return an array of Items from the loot filter file.
//
// The format appears to be three columns (ID, Icon ID, and name) separated by
// a "^" character. The first line is a header line.
//
// This will do some simple checking of the file format and if any issues
// are identified, it will return an error instead of any data.
//
// Error reasons:
//   - File cannot be opened for reading.
//   - Header line does not match expected header.
//   - Column 1 or 2 are not integers.
func ReadLF(fname string) ([]LFItem, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Check header line
	_ = scanner.Scan()
	head := scanner.Text()
	if head != lfHeader {
		// Unexpected header line. Either not a loot filter file
		// or format may have changed.
		return nil, fmt.Errorf("%s: missing expected header of \"%s\"",
			fname, lfHeader)
	}
	lineNo := 1
	var items []LFItem
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		parts := strings.Split(line, "^")
		if len(parts) != 3 {
			return nil, fmt.Errorf("%s at line %d has %d columns instead of 3",
				fname, lineNo, len(parts))
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 1 is not an integer",
				fname, lineNo)
		}
		iconID, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 2 is not an integer",
				fname, lineNo)
		}
		items = append(items, LFItem{
			ID:     id,
			IconID: iconID,
			Name:   parts[2],
		})
	}
	return items, err
}
