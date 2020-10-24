# Downloading and installation <!-- omit in TOC -->

- [1. Overview](#1-overview)
- [2. Downloading Supporting Data Files](#2-downloading-supporting-data-files)
  - [2.1. Zip file containing the data files](#21-zip-file-containing-the-data-files)
  - [2.2. Git clone](#22-git-clone)
- [3. Getting the programs](#3-getting-the-programs)
  - [3.1. Downloading the binary](#31-downloading-the-binary)
  - [3.2. Downloading/compiling the source](#32-downloadingcompiling-the-source)
    - [3.2.1. Zip file](#321-zip-file)
    - [3.2.2. Git clone](#322-git-clone)

## 1. Overview
These programs are available in two forms; binary and source. There are also two data
files that should be downloaded to get started.

## 2. Downloading Supporting Data Files

The two data files are the collection_quests.yml and the itemdb.yml. Two methods follow.

### 2.1. Zip file containing the data files

1) Visit the GitHub repo at "https://github.com/nuttann/eqdata".
2) Click on the button labeled "Code" that has a down arrow on it and select
   "Download ZIP"
3) Extract the files from the downloaded ZIP file.

### 2.2. Git clone

One advantage of "git clone" is that you can easily get new versions of the
files if they are updated on GitHub.  This would be the best way if I was still
playing and keeping the info current. Since I have stopped playing, this may or
may not happen. I may if someone I know passes me info on new quests and
inventory or real estate files with newer items in them.

1) Have "git" installed. (Check the web for how to do this.)
2) Type "git clone https://github.com/nuttann/eqdata.git" in a terminal window
   where you want the "eqdata" directory to be created.
3) If updates are made on the server "git pull" from inside the "eqdata"
   directory should get the new versions.

## 3. Getting the programs

### 3.1. Downloading the binary

1) Visit TODO
2) Download TODO

### 3.2. Downloading/compiling the source

If you want to download and build from the source, there are two ways to do it.
Zip gives you a snapshot of the source. Git will make it easier to get updates.

#### 3.2.1. Zip file

#### 3.2.2. Git clone

source whether just to understand, to compile, or to modify.

1) Have "git" installed. (Check the web for how to do this.)
2) Type "git clone https://github.com/nuttann/eqtools.git" in a terminal window
   where you want the source directory to be created.
3) Have "go" installed. See "https://golang.org/doc/install" for instructions.
4) TODO - build/install

Example download and build

- "cd"
- "mkdir Eq"
- "cd Eq"
- "git clone https://github.com/nuttann/eqtools.git"
- "cd eqtools"