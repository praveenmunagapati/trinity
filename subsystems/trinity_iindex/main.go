package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/clownpriest/trinity/core/commands"
	"github.com/clownpriest/trinity/core/config"
)

var iindex IIndex
var mainWaitGroup sync.WaitGroup

func main() {
	if len(os.Args) != 2 {
		fmt.Println(commands.IIndexHelp)
		os.Exit(0)
	}

	config, err := config.LoadIIndexConfig(os.Args[1])
	if err != nil {
		panic(err)
	}
	var store *IIndexStore
	store, err = NewIIndexStore(config)
	if err != nil {
		panic(err)
	}

	iindex = NewIIndex(store, config)

	mainWaitGroup.Add(1)

	go startIIndexServer(iindex.config.GRPCPort, &mainWaitGroup)

	mainWaitGroup.Wait()

}
