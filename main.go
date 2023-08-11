package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exec() error {
	ctx, cancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFn()

	opts, err := getOptions()
	if err != nil || opts == nil {
		// both opts and err will be nil, if help flag specified
		return err
	}

	// print version and exit
	if opts.PrintVersion {
		fmt.Println(version(opts.LogLevel() != 0))
		return nil
	}

	<-ctx.Done()

	return nil
}
