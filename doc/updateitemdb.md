# The "updateitemdb" command. <!-- omit in TOC -->

- [1. Overview](#1-overview)
- [2. Features](#2-features)
- [3. Limitations](#3-limitations)
- [4. Usage](#4-usage)
- [5. Configuration file format](#5-configuration-file-format)
  - [5.1. itemdbloc](#51-itemdbloc)
  - [5.2. realestate](#52-realestate)
  - [5.3. inventories](#53-inventories)
  - [5.4. lootfilters](#54-lootfilters)
- [6. Downloading and installation](#6-downloading-and-installation)

## 1. Overview

The "updateitemdb" command is capable of reading in certain types of EQ data
sources to update the item DB. This item DB is used by the "collectstoweb" program to display quest item names that are not in a house.

- Inventory files produced with "/output inventory" from within EQ.
- Real-estate files produced with "/output realestate" from within EQ.
- Loot filter files produced by EQ when using the loot filter UI from within
  EQ.

This was written to extract this data from local files for later use rather than
beat on servers like Allakhazam to scrape the data from as needed.

## 2. Features

- Real-estate files can be configured to be used.
- Inventory files can be configured to be used.
- Loot filter files can be configured to be used.
- Item names from all configured files **can** be used to update the item DB.
- Item icon IDs will be read from loot-filter files to update the item DB.

## 3. Limitations

These limitations are mainly due to not being able to know which data in files
was current. Some improvement can be made, but the loot-filter files
specifically contains a lot of outdated information even if the file has been
modified recently. These limitations are what made sense in my case. Inventory
and real-estate files that I had EQ create were done about the same time.
Loot-filter files were taken with a grain of salt.

- If older inventory and real-estate files are configured to be read, old names
  could be set.
- Names from loot-filter are only used if the name was not previously set.

## 4. Usage

The usage is very simple once the [configuration](#configuration-file-format)
is set up properly.

The program has a "conf" argument that must be present and point to the
configuration file to use. If the program is installed in your PATH, the
command is simply the following:

updateitemdb -conf PATH-TO-CONFIG-FILE

Example:

updateitemdb -conf /Users/Nuttann/Eq/conf/iteminfo_conf.yml

As usual, use quotes where necessary if the paths used have spaces.  You can
also set up a desktop shortcut to do this. If you use a shortcut, set it so
that the window does not disappear automatically when done.  Otherwise, you
will miss any messages that are printed. 

## 5. Configuration file format

See the configuration file in "samples/iteminfo_conf.yml" for an example.
There are brief comments in that file and sample entries. A detailed
description is given here.

### 5.1. itemdbloc

This "Item Database Location" should point to the location of the database. The
current code only supports a YAML file as the database. The amount of data is
small enough that the YAML version works fine. See [Downloading and
Installation](./downloading.md) for more information this file.

### 5.2. realestate

This parameter is a list of paths/filenames to real-estate dumps created by
"/output realestate" for a character. These are files that by default are
placed in the EQ install directory with names that follow the pattern
"CHARACTER_SERVER-RealEstate.txt" (E.g., Nuttann_cazic-RealEstate.txt").

### 5.3. inventories

This parameter is a list of paths/filenames to inventory dumps created by
"/output inventory" for a character. These are files that by default are placed
in the EQ install directory with names that follow the pattern
"CHARACTER_SERVER-Inventory.txt" (E.g., Nuttann_cazic-Inventory.txt).

### 5.4. lootfilters

This parameter is a list of paths/filenames to the various loot-filter files. These are files
that are updated when using the loot-filter UI and setting or removing a persistent setting (I.e., "Nvr" - Never loot, "AN" - Always Need, "AG" - Always Greed, and "Rnd" - Random). The
names are currently placed in the "userdata" folder of the EQ install directory and whose names are formatted like "LF_TYPE_CHARACTER_SERVER.ini" (E.g.,LF_AG_Nuttann_cazic.ini). TYPE is the persistent setting type mentioned above.

## 6. Downloading and installation

See the [Downloading and Installation](./downloading.md) document for instructions.
