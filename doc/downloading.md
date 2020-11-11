# Downloading and Installation <!-- omit in TOC -->

- [1. Overview](#1-overview)
- [2. Downloading Supporting Data Files](#2-downloading-supporting-data-files)
  - [2.1. Git clone of data files](#21-git-clone-of-data-files)
  - [2.2. Zip file of data files](#22-zip-file-of-data-files)
- [3. Getting the programs](#3-getting-the-programs)
  - [3.1. Download pre-compiled binary](#31-download-pre-compiled-binary)
  - [3.2. Download source code](#32-download-source-code)
    - [3.2.1. Git clone](#321-git-clone)
    - [3.2.2. Zip file](#322-zip-file)
  - [3.3 Building and installing from the source](#33-building-and-installing-from-the-source)

## 1. Overview
These programs are available in two forms; binary and source. There are also two data
files that should be downloaded to get started.

## 2. Downloading Supporting Data Files

The two data files are the collection_quests.yml and the itemdb.yml. The files
can be obtained using "git" or as a "zip" snapshot.

### 2.1. Git clone of data files

1) Have "git" installed. (Check the web for how to do this.)
2) Type "git clone https://github.com/nuttann/eqdata.git" in a terminal window
   where you want the "eqdata" directory to be created.
3) If updates are made on the server "git pull" from inside the "eqdata"
   directory should get the new versions.

### 2.2. Zip file of data files

1) Visit the GitHub repo at "https://github.com/nuttann/eqdata".
2) Click on the green button labeled "Code" that has a down arrow on it.
3) Click on "Download ZIP" selection in the dropdown.
4) Extract the files from the downloaded ZIP file.

## 3. Getting the programs

### 3.1. Download pre-compiled binary

1) Visit the GitHub repo at "https://github.com/nuttann/equtils".
2) Click on a release (usually latest) on the right side of the page.
3) Right click on the binary that you want and save to your system. (Chrome gave
   me a warning that it wasn't commonly downloaded so may be dangerous. I doubt
   it will get downloaded enough to ever commonly downloaded. There is always
   the source version below if you want to look at the source or compile from
   source.)
4) Place someplace so that you can run it once you have set up the proper configuration files and obtained the [Supporting Data Files](#2-downloading-supporting-data-files).

### 3.2. Download source code

If you want to build from source or just to understand it, there are two ways
to get the source. You can use "git" to clone the repository or simply obtain
a "zip" snapshot of the source.

#### 3.2.1. Git clone

1) Have "git" installed. (Check the web for how to do this.)
2) Type "git clone https://github.com/nuttann/eqtools.git" in a terminal window
   where you want the source directory to be created.
3) If updates are made on the server "git pull" from inside the "equtils"
   directory should get the new versions.

#### 3.2.2. Zip file

1) Visit the GitHub repo at "https://github.com/nuttann/equtils".
2) Click on the green button labeled "Code" that has a down arrow on it.
3) Click on "Download ZIP" selection in the dropdown.
4) Extract the files from the downloaded ZIP file.

### 3.3 Building and installing from the source

1) Have "go" installed. See "https://golang.org/doc/install" for instructions.
2) Download source as described above [Download source](#32-source-code).
3) "cd" into the source directory.
4) "go install ./..." - This will build all executables and place the results in
   go's binary directory. (Often "go/bin" in the user's home directory.)