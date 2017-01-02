package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/clownpriest/trinity/core/commands"
	"github.com/clownpriest/trinity/core/config"
)

var findex FIndex

var mainWaitGroup sync.WaitGroup

func main() {
	if len(os.Args) != 2 {
		fmt.Println(commands.FIndexHelp)
		os.Exit(1)
	}

	config, err := config.LoadFIndexConfig(os.Args[1])
	if err != nil {
		panic(err)
	}
	var store *FIndexStore
	store, err = NewFIndexStore(config)
	if err != nil {
		panic(err)
	}

	findex = NewFIndex(store, config)

	mainWaitGroup.Add(1)
	go startFIndexServer(findex.config.GRPCPort, &mainWaitGroup)

	mainWaitGroup.Wait()
}
