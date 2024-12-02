package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("No arguments given, nothing to do.\n")
		os.Exit(1)
	}

	subcmd := args[0]
	switch subcmd {
	case "-h", "-help", "help":
		fmt.Println("Run one of the sub commands with -h for more information:")
		fmt.Println("  run - Runs a given solution (make sure docker-compose is running")
	case "run", "r":
		runCmd(args[1:])
	default:
		// Default to run command, but since "run" wasn't mentioned we'll treat the
		// first argument as part of the invocation instead of stripping it.
		runCmd(args)
	}
}
