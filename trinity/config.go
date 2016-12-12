package main

import (
	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/system"
)

/*
LoadConfigs loads the root and subsystem configs into the
main configuration struct
*/
func LoadConfigs() (config.Config, error) {
	rootConfig, rootConfigErr := config.LoadRootConfig(system.RootConfigPath())
	if rootConfigErr != nil {
		return rootConfig, rootConfigErr
	}

	findexConfig, findexConfigErr := config.LoadFIndexConfig(system.FIndexConfigPath())
	if findexConfigErr != nil {
		return rootConfig, findexConfigErr
	}
	rootConfig.FIndex = findexConfig

	iindexConfig, iindexConfigErr := config.LoadIIndexConfig(system.IIndexConfigPath())
	if iindexConfigErr != nil {
		return rootConfig, iindexConfigErr
	}
	rootConfig.IIndex = iindexConfig

	return rootConfig, nil
}
