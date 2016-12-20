package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/clownpriest/trinity/core"
	"github.com/clownpriest/trinity/core/commands"
	"github.com/clownpriest/trinity/core/commands/cli"
	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/system"
)

// GitHash is the git HEAD commit hash the binary
// was built with
var GitHash string

var interuptHandler *system.InteruptHandler

var node *core.TrinityNode

var mainConfig config.Config

func main() {

	if len(os.Args) < 2 {
		fmt.Println(commands.RootHelp)
		os.Exit(0)
	}

	commandResult, commandErr := cli.Parse(os.Args)
	if commandErr != nil {
		printUsage()
		os.Exit(0)
	}

	handleCommand(commandResult)

	var err error
	node, err = LoadNode(&mainConfig)
	if err != nil {
		panic(err)
	}

	var mainWaitGroup sync.WaitGroup

	mainWaitGroup.Add(1)
	go startTrinityCoreServer(mainConfig.CoreServerPort, &mainWaitGroup)

	err = node.InitSubsystems()
	if err != nil {
		panic(err)
	}

	mainWaitGroup.Add(1)

	sigChan := make(chan os.Signal, 1)
	go startInterruptHandler(sigChan, &mainWaitGroup)

	mainWaitGroup.Wait()

}
