package main

import (
	"os"

	"github.com/clownpriest/trinity/core/config"
)

var iindexConfig *config.IIndexConfig

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(0)
	}

	iiConfig, err := config.LoadIIndexConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	iindexConfig = &iiConfig

}
