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

type Options struct {
	InputFile    string `short:"i" long:"input"   description:"name of the DayOne2 zip archive file"`
	OutputFolder string `short:"o" long:"output"  description:"output directory for the markdown files"`
	JournalName  string `short:"j" long:"journal" description:"name of JSON file in zip archive"`
	PrintVersion bool   `long:"version"           description:"print version and exit"`
	Verbose      []bool `short:"v" long:"verbose" description:"show verbose output, list multiple times for even more verbose output"`
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
