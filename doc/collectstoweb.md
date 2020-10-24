# The "collectstoweb" command. <!-- omit in TOC -->

- [1. Overview](#1-overview)
- [2. Features](#2-features)
- [3. Limitations](#3-limitations)
- [4. Usage](#4-usage)
- [5. Configuration file format](#5-configuration-file-format)
  - [5.1. questsfile](#51-questsfile)
  - [5.2. itemdbloc](#52-itemdbloc)
  - [5.3. htmlout](#53-htmlout)
  - [5.4. htmltitle](#54-htmltitle)
  - [5.5. htmlintro](#55-htmlintro)
  - [5.6. houses](#56-houses)
- [6. Future enhancements](#6-future-enhancements)
- [7. Downloading and installation](#7-downloading-and-installation)

## 1. Overview

The "collectstoweb" command is designed to assist in managing collection
items stored in houses. This takes the unorganized real estate dump and creates
an HTML page with the counts organized by a hierarchy of expansion, zone and
quest. This can be useful if offering these up to others such as a guild,
having joint management among a group of friends, or even just for personal use
by alts.

The implemented features were the ones that were useful in my specific case.  I
had gathered collection quest items over the years and made them available to
guild members upon request. Collection items are non-placeable, so can only be
**stored** in a house. At the time this was first written, the largest house
could only hold 300 non-placeable items. While a couple special crates could
hold an additional 100 or 200 non-placeable items, these were ignored so there
was only one place to store in a house. This made putting things in simpler.

It supports a note for each quest. This was useful for giving hints as to where
in a zone or what mobs dropped items for that quest. Some people never
requested items, but did bookmark the posted HTML just for quick info on where
to farm them. The program is data-driven so can be configured as to which
houses hold which quests. The note mentioned above is also from a data-driven
file. A customized title and additional information can be added to the HTML
page via the main configuration file.

In addition to building the HTML page, this will print out messages when
multiple stacks of an item are in a house or if items other than the ones for
the configured quests are stored in the house. There are a limited number of
stacks that can be stored in a house. Multiple stacks could exist for two
reasons. First, if you have over 100 of an item, a second stack would be made.
The second reason is that if multiple people are placing items into the house,
it could end up with partial stacks tagged with two different owners.
Personally, I dropped anything over one stack on the ground.  Also, I used to
allow some officers in a previous guild to organize stuff in the house and the
multiple stack method let me know that I had stacks to combine.

## 2. Features

- Houses can be configured as holding quest items for an expansion.
- Houses can be configured as holding quest items for zones in an expansion.
  (This was supported as Rain of Fear has more than 300 collection items.)
- Title and information to be added at the beginning of the output can be
  configured.
- Reports when multiple stacks of a collection item are in a house.
- Reports when there are stored items in a house that are not from the
  configured collections.

## 3. Limitations

Some limitations here are due to just meeting my use case.

- This does not support splitting quests from a zone into multiple houses. The
  same zone can be configured for multiple houses but all quests for that
  zone will be listed for both houses. This would have been more complicated
  code and it was easy enough to have alts get more houses.
- The current code only mentions if extra items were stored in a house. It
  doesn't state whether they are collection items, and if so, in which house
  they should have been.  There was an older Python program that went through
  my inventory dump and found collection items and would print out where they
  were in my bags and in which house they should be placed. I may get around to
  reimplementing that.
- Names that have not been collected for items show up as "???" for their name.
  Items that are in the houses will have a name along with those that have been
  seen previously or have been collected with the
  ["updateitemdb"](./updateitemdb.md) program.

## 4. Usage

The usage is very simple once the [configuration](#configuration-file-format)
is set up properly.

The program has a "conf" argument that must be present and point to the
configuration file to use. If the program is installed in your PATH, the
command is simply the following:

collectstoweb -conf PATH-TO-CONFIG-FILE

Example:

collectiontohtml -conf /Users/Nuttann/Eq/conf/collections_conf.yml

As usual, use quotes where necessary if the paths used have spaces.  You can
also set up a desktop shortcut to do this. If you use a shortcut, set it so
that the window does not disappear automatically when done.  Otherwise, you
will miss any messages that are printed. You may also want multiple
configuration files to switch between such as when you are doing this on
multiple servers.

## 5. Configuration file format

See the configuration file in "samples/collection_conf.yml" for an example.
There are brief comments in that file and sample entries. A detailed
description is given here.

### 5.1. questsfile

This should point to the file containing all the quest info. See [Downloading and
Installation](./downloading.md) for more information this file.

### 5.2. itemdbloc

This "Item Database Location" should point to the location of the database. The
current code only supports a YAML file as the database. The amount of data is
small enough that the YAML version works fine. See [Downloading and
Installation](./downloading.md) for more information this file.

### 5.3. htmlout

This parameter indicates the path of where to write the output file. This file
will be created if it doesn't exist and will be overwritten if it does exist.

### 5.4. htmltitle

This parameter holds the text to display as the page title and also the main
header. HTML-specific characters such as '&' will be escaped, so text here
should appear verbatim in a browser.

### 5.5. htmlintro

This parameter holds the **RAW** HTML to place into the final HTML document
verbatim after the HTML title/header and before the listing of the house
contents. This is to support any HTML markup or links to any instructions or
such that should be included. I had used a link with a sample request and
instructions to explain why the information to put in a request was really
useful to me when pulling things out.

### 5.6. houses

This parameter is the most complex of the configuration parameters. It is a
nested YAML definition. YAML uses indenting for the nesting and special syntax
for definitions and lists. There are times when double quotes will be necessary
or YAML will not parse the data as the user intended. The sample file has a
good example of some of the necessary things. If the parser complains and the
indentation looks fine, try adding double quotes around the definition. It may
contain a special character that you are unaware of.

You can search on the web for YAML specifications.  The following is a summary
of stuff that exists in the sample.

- Key-value pairs are defined by "key: value" there is a required space after
  the ':'.
- Indentation defines nesting.
- A "- " at the beginning indicates an array element.
- A "#" at the beginning of the line or after whitespace indicates the rest of
  the line is a comment.

The basic nesting is as follows:

```YAML
# houses is the root element of the housing configuration. It holds an array
# of house configurations.
houses:
  # The '-' before fname means this is an array element. The first field in the
  # array is fname, property is indented the same as fname, so is the second field,
  # and contents is also indented the same so is the 3rd field.
  - fname: "C:/Users/Public/Daybreak Game Company/Installed Games/Everquest/Nuttann_cazic-RealEstate.txt"
    property: "Return of the Exiled Village II, 113 Vanward Street, Bixie Hive House"
    expansions: # This case is the same as above. The value below is an array of expansions.
      # Each expansion has a name and an OPTIONAL array of zones.
      - name: Rain of Fear
        # The zones attribute exists for this expansion, so only quests from those zones
        # are included.
        zones:
          - Multi-zone
          - Shard's Landing
          - East Wastes, Zeixshi-Kar's Awakening
          - Kael Drakkel, The King's Madness
          - Crystal Caverns, Fragment of Fear
          - Evantil, the Vile Oak
          - Breeding Grounds
          - Chapterhouse of the Fallen
          - Grelleth's Palace, Chateau of Filth

  - fname: "C:/Users/Public/Daybreak Game Company/Installed Games/Everquest/Nuttann_cazic-RealEstate.txt"
    address: "Return of the Exiled Village II, 111 Vanward Street, Evantil's Abode"
    expansions: # This house has two expansions (RoF and the pseudo-expansion Periodic)
      - name: Rain of Fear
        zones:
          - Valley of King Xorbb
          - Chelsith Reborn
          - Plane of Shadow
          - Heart of Fear
      - name: Periodic 
        # No zones listed, so all zones for this "pseudo-expansion" are included.

  - fname: "C:/Users/Public/Daybreak Game Company/Installed Games/Everquest/Nuttann_cazic-RealEstate.txt"
    address: "Return of the Exiled Village II, 112 Vanward Street, Bixie Hive House"
    contents:
      - name: Call of the Forsaken
        # No zones listed, so all CotF zones are included.
```

## 6. Future enhancements

Many of the limitations were due to just meeting my personal needs. There are
some things that could be done to make this more useful.

Possible future work:

- Report whether stored items in a house that don't belong to the configured
  quests to that house are colection items or not.
- If a collection item is in a wrong house, report which house it should be in.
- Add feature to search bags for collection items and list which house they
  should should be place in. (An older Python program I wrote used to do this
  for me and it made it quicker to drop items into the proper house.)

## 7. Downloading and installation

See the [Downloading and Installation](./downloading.md) document for instructions.
