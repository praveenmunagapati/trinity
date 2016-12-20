package main

import (
	"os"

	"github.com/clownpriest/trinity/core/commands/cli"
	"github.com/clownpriest/trinity/core/system"
)

func handleCommand(com cli.CommandType) error {
	if com == cli.Initialize {
		err := system.Init()
		if err != nil {
			if err == system.ErrTrinityDirExists {
				os.Exit(0)
			}
		}
		os.Exit(0)

	} else if com == cli.Version {
		printVersion()
		os.Exit(0)

	} else if com == cli.Start {
		config, configErr := LoadConfigs()
		if configErr != nil {
			panic(configErr)
		}

		mainConfig = config

	}
	return nil
}
