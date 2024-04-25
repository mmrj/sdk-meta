package main

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/cmd/log/commands"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `Create and manage log specifiers
		USAGE:
			log <command>

		COMMANDS:
			new: Create a new system, class, or condition.
			deprecate: Deprecate a log code.
			supersede: Indicate a log code has been superseded.
			document: Generate markdown documentation for log codes.
			validate: Validate codes against the schema.
		`)
}

func validateCodes() {
	err := logs.ValidateCodes()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		validateCodes()
		commands.RunNewCommand()
	case "deprecate":
		validateCodes()
		commands.RunDeprecateCommand()
	case "supersede":
		validateCodes()
		commands.RunSupersedeCommand()
	case "document":
		validateCodes()
		commands.RunDocumentCommand()
	case "validate":
		validateCodes()
		fmt.Println("Codes matched schema.")
	default:
		fmt.Printf("Unrecognized command: %s\n", os.Args[1])
		usage()
	}

	os.Exit(0)
}
