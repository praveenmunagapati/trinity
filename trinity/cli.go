package main

import (
	"fmt"

	"github.com/clownpriest/trinity/core/commands"
)

func printVersion() {
	fmt.Printf("\ntrinity v" + version() + "\n\n")
}

func printUsage() {
	fmt.Println(commands.RootHelp)
}
