package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/clownpriest/trinity/core/config"
)

var (
	// ErrTrinityDirExists means the $HOME/.trinity directory already exists
	ErrTrinityDirExists = errors.New("init: trinity directory already exists")

	// ErrConfigFileAlreadyExists means that there's already a config file in $HOME/.trinity
	ErrConfigFileAlreadyExists = errors.New("init: config file already exists in $HOME/.trinity")
)

/*
Init sets up the $HOME/.trinity directory and
initializes the node's
*/
func Init() error {
	_, err := os.Stat(configsDir)
	if err == nil {
		fmt.Printf("\ntrinity node has already been initialized\n\n")
		return ErrTrinityDirExists
	}
	if os.IsNotExist(err) {
		createTrinityDir()

		err = createDatastoreDir()
		if err != nil {
			return err
		}

		err = createConfigsDir()
		if err != nil {
			return err
		}

		err = writeDefaultConfigs()
		if err != nil {
			return err
		}
	}
	fmt.Printf("\ntrinity node initialized\n\n")
	return nil
}

func createTrinityDir() error {
	err := os.Mkdir(trinityDir, os.ModePerm)
	return err
}

func createDatastoreDir() error {
	err := os.Mkdir(datastoreDir, os.ModePerm)
	return err
}

func createConfigsDir() error {
	err := os.Mkdir(configsDir, os.ModePerm)
	return err
}

func writeDefaultConfigs() error {
	rootConfig := config.NewDefaultConfig()
	gatewayConfig := config.NewDefaultGatewayConfig()
	findexConfig := config.NewDefaultFIndexConfig()
	iindexConfig := config.NewDefaultIIndexConfig()

	findexConfig.StorePath = FIndexPath()
	iindexConfig.StorePath = IIndexPath()

	rootConfig.Gateway = gatewayConfig
	rootConfig.FIndex = findexConfig
	rootConfig.IIndex = iindexConfig

	var err error

	err = writeConfig(rootConfigPath, rootConfig)
	if err != nil {
		return err
	}

	err = writeConfig(gatewayConfigPath, gatewayConfig)
	if err != nil {
		return err
	}

	err = writeConfig(tfindexConfigPath, findexConfig)
	if err != nil {
		return err
	}

	err = writeConfig(tiindexConfigPath, iindexConfig)
	if err != nil {
		return err
	}

	return nil
}

func writeConfig(configPath string, config interface{}) error {
	_, gatewayErr := os.Stat(configPath)
	if gatewayErr == nil {
		fmt.Println("config file already exists in $HOME/.trinity/configs")
		return ErrConfigFileAlreadyExists
	}

	if os.IsNotExist(gatewayErr) {
		file, createErr := os.Create(configPath)
		defer file.Close()
		if createErr == nil {
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			encoder.Encode(config)
		}
	}
	return nil
}
