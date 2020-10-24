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

// REItem is the data found in each line of a RealEstate file.
type REItem struct {
	RELoc    string // Neighborhood = yard, Plot = house
	REName   string // Address, address and house type
	ItemName string // EQ name
	Owner    string // Toon owning item
	Status   string // E.g., Placed, Stored
	ID       int    // EQ ident key
	Count    int    // Number of that item
}

// Expected real-estate header line
const reHeader = "RealEstateLocation\tRealEstateName\tItemName\tItemOwner\tStatus\tID\tCount"

// ReadRE will return an array of Items from the real-estate file.
//
// The format appears to be seven columns (RE location, RE name, Item name,
// Item owner, Status, ID, and Count) separated by a tab character. The first
// line is a header line.
//
// This will do some simple checking of the file format and if any issues are
// identified, it will return an error instead of any data.
//
// Error reasons:
//   - File cannot be opened for reading.
//   - Header line does not match expected header.
//   - Column 6 or 7 are not integers.
func ReadRE(fname string) ([]REItem, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Check header line
	_ = scanner.Scan()
	head := scanner.Text()
	if head != reHeader {
		// Unexpected header line. Either not a loot filter file or format may
		// have changed.
		return nil, fmt.Errorf(
			"%s: first line does not match expected header of \"%s\"",
			fname, reHeader)
	}
	lineNo := 1
	var items []REItem
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		parts := strings.Split(line, "\t")
		if len(parts) != 7 {
			return nil, fmt.Errorf("%s at line %d has %d columns instead of 3",
				fname, lineNo, len(parts))
		}
		id, err := strconv.Atoi(parts[5])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 6 is not an integer",
				fname, lineNo)
		}
		count, err := strconv.Atoi(parts[6])
		if err != nil {
			return nil, fmt.Errorf("%s at line %d column 7 is not an integer",
				fname, lineNo)
		}
		items = append(items, REItem{
			RELoc:    parts[0],
			REName:   parts[1],
			ItemName: parts[2],
			Owner:    parts[3],
			Status:   parts[4],
			ID:       id,
			Count:    count,
		})
	}
	return items, nil
}
