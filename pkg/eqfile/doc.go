// Copyright 2020 Nuttann. All rights reserved.
// The use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

/*
Package eqfile provides functions for interacting with EQ produced files.

The various "Read..." functions will read and return the data from files. The
formats of these files may change over time. This package is intended to
provide an abstraction layer to allow changes in the files to  be handled in
one place.

These routines will return an error instead of data if anything doesn't match
what is expected. This is intended to help detect issues with the files in case
the file is of the wrong type or if the EQ file format has changed. This
checking is not foolproof. The first check is to see if the header line matches
exactly the expected one for that type. Ideally, the developers of EQ will not
forget to update the header line if they changed the format. The additional
checks per line are that expected columns of integers are integers and that
each line has the expected number of columns.

Supported file types are EQ files such as those produced by EQ. This currently
includes the following file types:

  - RealEstate
  - Inventory
  - Loot Filter

File formats - I did not find any documented format so determined the formats
by inspection. They each have one header line followed by data lines. Each data
line consists of columns separated by a character. These do not parse with CSV
libraries properly when quotes appear in a field as they are not escaped per
the CSV specification. Instead, a simple splitting on the special character
appears to work and so far, that character does seem to appear within a field.

Real Estate

Real estate files are created by the "/output realestate [FILENAME]" command in
EQ. They contain content listings for all real estate owned by the user
executing the command. All data in these files should be valid at the time the
command was executed. The default file name is of the form
"{TOON}_{SERVER}-RealEstate.txt" and found in the EQ install directory. Users
may add an optional file name at the end of the command to direct the output to
a different file.

Inventory

Inventory files are created by the "/output inventory [FILENAME]" command in
EQ. They contain listings for all items in the user's inventory and bank,
including the shared bank slots. All data in these files should be valid at the
time the command was executed. The default file name is of the form
"{TOON}_{SERVER}-Inventory.txt" and found in the EQ install directory. Users
may add an optional file name at the end of the command to direct the output to
a different file.

Loot Filters

Loot filter files are created and updated automatcally when using the advanced
loot interface in EQ and selecting any of the retained settings of Always Need
(AN), Always Greed (AG), Never (Nvr), or Random (Rnd). These are stored in four
files, one for each retained setting of the form LF_{TYPE}_{TOON}_{SERVER}.ini
where TYPE is one of (AN - Always Need, AG - Always Greed, Nvr - Never, or Rnd
- Random). These files contain data for the last retained setting of the item.
Only items being updated at the time are updated in the file, so many entries
may have old names.
*/
package eqfile
