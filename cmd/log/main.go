package main

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/cmd/log/commands"
)

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `Create and manage log specifiers
		USAGE:
			log <command>

		GET HELP FOR A COMMAND:
			log <command> --help

		COMMANDS:
			new: Create a new system, class, or condition.
			deprecate: Deprecate a log code.
			supersede: Indicate a log code has been superseded.
		`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		commands.RunNewCommand()
	case "deprecate":
		// TODO
	case "supersede":
		// TODO
	default:
		fmt.Printf("Unrecognized command: %s\n", os.Args[1])
		usage()
	}

	os.Exit(0)
}
