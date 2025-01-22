package app

import (
	"fmt"
	"gosolve/internal/cli"
	"gosolve/internal/version"
	"os"
)

// Main is the main entry point of the application.
// It checks if the first command line argument is "--version" and if so, prints
// the version of the application and exits.
// If not, it configures the application logger and starts the command line
// interface.
func Main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("GoSolveAPI version", version.Version())
		os.Exit(0)
	}

	// Starting command line interface
	cli.NewRootCmd().Execute()
}
