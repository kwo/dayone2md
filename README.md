# Dayone To Markdown

## TODO
* https://github.com/ncruces/go-sqlite3

Export Dayone to Markdown.

[![tag](https://img.shields.io/github/tag/kwo/dayone2md.svg)](https://github.com/kwo/dayone2md/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/kwo/dayone2md?status.svg)](https://pkg.go.dev/github.com/kwo/dayone2md)
[![Go report](https://goreportcard.com/badge/github.com/kwo/dayone2md)](https://goreportcard.com/report/github.com/kwo/dayone2md)
[![Contributors](https://img.shields.io/github/contributors/kwo/dayone2md)](https://github.com/kwo/dayone2md/graphs/contributors)
[![License](https://img.shields.io/github/license/kwo/dayone2md)](./LICENSE)


`dayone2md` is a CLI application that can create a directory a markdown files from entries in [Dayone](https://dayoneapp.com/).
Use either a Dayone JSON archive or read directly from the Dayone database.
Currently only MacOS is supported.


## Features

- Input sources:
  - JSON export archive
  - read directly from the Dayone database
- Photos: download embedded photos and rewrite photo links
- Wikilinks: rewrite links to other entries as wikilinks
- Templates: supports external templates
- Title: separate out first line of posts as a title to format as desired
- In-place updates to output folder
- Conditional updates: only write files to output folder that differ from source
- Option to keep files in the output folder that do not exist at the source
- Flag to group posts by day in a single file
- Flag to alter sort order within day when grouped by day
- Installation: static binary without dependencies on external libraries


## 🚀 Install

```sh
brew install kwo/tools/dayone2md
```

## 💡 Usage

```
Usage:
  dayone2md [OPTIONS]

Application Options:
  -j, --journal=      journal name to export
  -i, --input=        input file, either the DayOne.sqlite database file or the JSON export zip file
  -o, --output=       output directory
  -t, --template=     name of the template to use, either the path of an external template file or the name of a built-in template: main or
                      full (default: main)
  -g, --group         group entries by day, one file per day, multiple entries per file
  -r, --reverse       reverse chronological sort order for entries within a file, useful only if entries are grouped by day
      --keep-orphans  do not remove files in the output directory that lack a matching entry in the input file
      --version       print version and exit
  -v, --verbose       show verbose output, list multiple times for even more verbose output

Help Options:
  -h, --help          Show this help message
```

### Example Usage

```sh
dayone2md -i "$HOME/Library/Group Containers/5U8NS4GX82.dayoneapp2/Data/Documents/DayOne.sqlite" -o $HOME/Documents/Journal -j Journal -g -vv

OR

dayone2md -i "$HOME/Desktop/09-07-2023_8-30-PM.zip" -o $HOME/Documents/Journal -j Journal -g -vv
```

## Input Sources

### JSON Export archive

Follow these [instructions](https://dayoneapp.com/guides/tips-and-tutorials/exporting-entries/#mac) to export entries from the Dayone app to a JSON archive.

Then, use the archive location as the import source, for example:

`dayone2md -i "$HOME/Desktop/09-07-2023_8-30-PM.zip" ...`

### Database

The Dayone database on a Mac is located at

`$HOME/Library/Group Containers/5U8NS4GX82.dayoneapp2/Data/Documents/DayOne.sqlite`

To export entries directly from the database, use this location as the import source, like this:

`dayone2md -i "$HOME/Library/Group Containers/5U8NS4GX82.dayoneapp2/Data/Documents/DayOne.sqlite" ...`

Day One occasionally renames internal join tables (for example `Z_13TAGS`, `Z_42TAGS`, etc.). `dayone2md` auto-detects the correct table and column names at runtime, so tag exports continue to work regardless of the specific schema version.

The database is opened in readonly mode and an export can be performed while the Dayone app is open. Also, note that for premium users that are synchronizing data from multiple devices, the Dayone app must be run once before exporting to retrieve entries from the Dayone cloud service.

Disclaimer: be careful by keeping a backup of your data. As stated in the license, this software assumes no liability for data loss.

## 🧑‍💻 development

## go-releaser

once:
```sh
brew install goreleaser golangci-lint
go install -v golang.org/x/vuln/cmd/govulncheck@latest
```

ongoing:
```shell
goreleaser release --clean --snapshot
./dist/dayone2md_darwin_arm64/dayone2md --version -vv
```

## similar projects
* [joshuacoles/Dayone-Export](https://github.com/joshuacoles/Dayone-Export)
* [quantumgardener/dayone-to-obsidian](https://github.com/quantumgardener/dayone-to-obsidian)
