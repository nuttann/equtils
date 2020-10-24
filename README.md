# EQ Utilities <!-- omit in TOC -->

- [1. Overview](#1-overview)
- [2. Programs](#2-programs)
  - [2.1. "collectstoweb"](#21-collectstoweb)
  - [2.2. "updateitemdb"](#22-updateitemdb)
- [3. Library packages ("go" source only)](#3-library-packages-go-source-only)
  - [3.1. eqfile](#31-eqfile)
  - [3.2. eqdb](#32-eqdb)
- [4. Downloading](#4-downloading)

## 1. Overview

This repository includes some programs to manage and display EQ data and are
available in both binary and source form. The source form also includes some
"go" packages to interface with EQ-produced data files and the extracted data.

## 2. Programs

### 2.1. "collectstoweb" 

Retrieve data from one or more real-estate dump files and create a web page
that organizes the collection items inventory by expansion / zone / quest. This
can be used for advertising available items if making them available to
guildies, a fellowship, other friends, or even just your alts. The organization
is a much easier presentation to check for needed items that are often spread
across many houses due to the sheer number of items.

See the [Detailed Documentation](./doc/collectstoweb.md) for usage
including configuration and examples.

### 2.2. "updateitemdb"

Read various EQ sources to update the EQ item DB. This is to help get item info
such as names and icon IDs from EQ output files and populate a database for use
by other programs.

See the [Detailed Documentation](./doc/updateitemdb.md) for usage including
configuration and examples.

## 3. Library packages ("go" source only)

### 3.1. eqfile

Read various EQ output files and provide an interface to get the data. This is
intended to isolate where changes would need to be made if file formats change
at some point.

- Real-estate dumps (I.e., "/output realestate" files)
- Inventory dumps (I.e., "/output inventory" files)
- Loot filter settings (I.e., LF_* files found in "userdata" folder in the EQ
  installtion folder)

### 3.2. eqdb

Store and retrieve data extracted from output files and other sources. This is
an abstraction layer intended to isolate the programs from the actual storage.
This could be changed to a real database or a network API if needed. It
currently uses YAML files for the data storage, which seems to supply
acceptable performance.

## 4. Downloading

[Downloading and Installation](./downloading.md)