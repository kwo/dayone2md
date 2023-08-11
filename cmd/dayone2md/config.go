package main

import (
	"errors"
	"fmt"

	"github.com/jessevdk/go-flags"
)

func getOptions() (*Options, error) {
	opts := &Options{}
	_, err := flags.Parse(opts)
	if err != nil {
		var e *flags.Error
		if errors.As(err, &e) && e.Type == flags.ErrHelp {
			return nil, nil // help requested
		}
		return nil, fmt.Errorf("cannot parse flags: %w", err)
	}
	return opts, nil
}

// see https://pkg.go.dev/github.com/jessevdk/go-flags?utm_source=godoc#hdr-Available_field_tags
type Options struct {
	JournalName     string `short:"j" long:"journal"                 description:"journal name to export"`
	InputLocation   string `short:"i" long:"input"                   description:"input file, either the DayOne.sqlite database file or the export zip file"`
	OutputDirectory string `short:"o" long:"output"                  description:"output directory"`
	Template        string `short:"t" long:"template" default:"main" description:"name of the template to use, either the path the external template file or the name of a built-in template: main or full"`
	GroupByDay      bool   `short:"g" long:"group"                   description:"group entries by day, one file per day, multiple entries per file"`
	SortReverse     bool   `short:"r" long:"reverse"                 description:"reverse chronological sort order for entries within a file"`
	KeepOrphans     bool   `          long:"keep-orphans"            description:"do not remove files at destination that lack a matching entry at the source"`
	PrintVersion    bool   `          long:"version"                 description:"print version and exit"`
	Verbose         []bool `short:"v" long:"verbose"                 description:"show verbose output, list multiple times for even more verbose output"`
}

func (o *Options) GetJournalName() string {
	return o.JournalName
}

func (o *Options) GetInputLocation() string {
	return o.InputLocation
}

func (o *Options) GetOutputDirectory() string {
	return o.OutputDirectory
}

func (o *Options) GetTemplate() string {
	return o.Template
}

func (o *Options) IsGroupByDay() bool {
	return o.GroupByDay
}

func (o *Options) IsSortReverse() bool {
	return o.SortReverse
}

func (o *Options) IsRemoveOrphans() bool {
	return !o.KeepOrphans
}

func (o *Options) LogLevel() int {
	level := 0
	for _, v := range o.Verbose {
		if v {
			level++
		}
	}
	return level
}
