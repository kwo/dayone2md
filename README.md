# DayOne To Markdown

Export DayOne databases to Markdown files.

## Usage

Export the archive from DayOne.

```
Usage:
  dayone2md [OPTIONS]

Application Options:
  -j, --journal=      journal name to export
  -i, --input=        input file, either the DayOne.sqlite database file or the export zip file
  -o, --output=       output directory
  -t, --template=     name of the template to use, either the path the external template file or the name of a built-in template: main or
                      full (default: main)
  -g, --group         group entries by day, one file per day, multiple entries per file
  -r, --reverse       reverse chronological sort order for entries within a file
      --keep-orphans  do not remove files at destination that lack a matching entry at the source
      --version       print version and exit
  -v, --verbose       show verbose output, list multiple times for even more verbose output

Help Options:
  -h, --help          Show this help message
```

## development

```shell
make lint-fix
make lint
make install
```

## similar projects
* [joshuacoles/Dayone-Export](https://github.com/joshuacoles/Dayone-Export)

## Production Ready (the last 80%)
 - [ ] update readme with features list
 - [ ] license
 - [ ] github repo
 - [ ] blog post, Twitter
 - [ ] publish doc to kwo.dev
 - [ ] goreleaser

## TODO
- [ ] abstract destination behind an interface, add memory impl for testing
- [ ] fill out database entity fields
- [x] attempt to use alternative sqlite lib
- [x] move cmd to package, move lib to day2onemd
- [x] rewrite dayone entry links
- [x] read directly from database
- [x] separate out first line of each entry as title
- [x] add option to alter sort order within day when grouped by day
- [x] add flag to group entries by day (for me the default)
- [x] add full date-time to full template frontmatter
- [x] exclude weather and location from metadata footer when they are not available
- [x] bugfix date in footer - looks like UTC not local
- [x] extract directly from zip without exploding first
- [x] multiple entries on same day
- [x] option to use own template
- [x] use 1.21 slog
- [x] remove orphans (--keep-orphans to prevent)
- [x] conditional write for photos and entries based on SHA hash
- [x] add location, weather, moonphase, etc to frontmatter
- [x] inplace update of iA directory
- [x] options to not add tags to bottom

