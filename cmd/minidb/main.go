package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const version = "0.1.0"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error: err")
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("minidb", flag.ContinueOnError)
	fs.SetOutput(stdout)
	showVersion := fs.Bool("version", false, "display version information")

	fs.Usage = func() {
		fmt.Fprintln(stdout, "MiniDB - SQLite-inspired db written in go")
		fmt.Fprintln(stdout, "\nUsage:")
		fmt.Fprintln(stdout, " minidb [options]")
		fmt.Fprintln(stdout, "\nOptions:")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected argument %s", fs.Arg(0))
	}
	if *showVersion {
		fmt.Fprintf(stdout, "MiniDB v%s\n", version)
		return nil
	}
	fs.Usage()
	return nil
}
