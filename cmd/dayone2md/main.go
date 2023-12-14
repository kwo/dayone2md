package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/lmittmann/tint"

	"github.com/kwo/dayone2md"
)

//nolint:gochecknoglobals
var (
	commithash = "unknown"
	commitTS   = "unknown"
	version    = "unknown"
)

func main() {
	verbosity := &slog.LevelVar{}
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      verbosity,
		TimeFormat: time.TimeOnly,
	}))
	slog.SetDefault(logger)

	ctx, cancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFn()

	opts, err := getOptions()
	if err != nil {
		// don't log, already printed - slog.ErrorContext(ctx, "cannot load options", "err", err)
		os.Exit(1)
	} else if opts == nil {
		return // help requests, just exit
	}

	// set log level
	switch opts.LogLevel() {
	case 0:
		verbosity.Set(slog.LevelError)
	case 1:
		verbosity.Set(slog.LevelWarn)
	case 2:
		verbosity.Set(slog.LevelInfo)
	default:
		verbosity.Set(slog.LevelDebug)
	}

	// print version and exit
	if opts.PrintVersion {
		if verbosity.Level() < slog.LevelWarn {
			ts, _ := strconv.ParseInt(commitTS, 10, 64)
			dt := time.Unix(ts, 0).UTC().Format(time.DateOnly)
			fmt.Printf("dayone2md %s %s %s %s\n", version, commithash, dt, runtime.Version())
		} else {
			fmt.Printf("dayone2md %s\n", version)
		}
		return
	}

	if opts.JournalName == "" || opts.InputLocation == "" || opts.OutputDirectory == "" {
		slog.ErrorContext(ctx, "please specify a journal (-j), and input location (-i), and an output directory (-o)")
		os.Exit(1)
	}

	// finally perform the conversion
	if err := dayone2md.Convert(ctx, opts); err != nil {
		slog.ErrorContext(ctx, "conversion failed", "err", err)
		os.Exit(1)
	}
}
