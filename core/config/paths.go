package config

import (
	"os"
	"path"
)

/*
These are the default paths to an assortment of system
files and directories.
*/
var (
	homeDir           = os.Getenv("HOME")
	trinityDir        = path.Join(homeDir, ".trinity")
	configsDir        = path.Join(trinityDir, "configs")
	datastoreDir      = path.Join(trinityDir, "datastore")
	subsystemDir      = path.Join(trinityDir, "subsystems")
	staticDir         = path.Join(trinityDir, "static")
	rootConfigPath    = path.Join(configsDir, "trinity_config")
	gatewayConfigPath = path.Join(configsDir, "gateway_config")
	tfindexConfigPath = path.Join(configsDir, "tfindex_config")
	tiindexConfigPath = path.Join(configsDir, "tiindex_config")
	tcrawlConfigPath  = path.Join(configsDir, "tcrawl_config")
	findexPath        = path.Join(datastoreDir, "findex")
	iindexPath        = path.Join(datastoreDir, "iindex")
)

// RootConfigPath returns the path to the main config file
func RootConfigPath() string {
	return rootConfigPath
}

// SubsystemsDir returns the directory with all the subsystem plugins
func SubsystemsDir() string {
	return subsystemDir
}

// StaticDir returns the directory with all the static content
func StaticDir() string {
	return staticDir
}

// FIndexConfigPath returns the path to the forward index config file
func FIndexConfigPath() string {
	return tfindexConfigPath
}

// IIndexConfigPath returns the path to the inverted index config file
func IIndexConfigPath() string {
	return tiindexConfigPath
}

// FIndexPath returns the path to the forward index datastore
func FIndexPath() string {
	return findexPath
}

// IIndexPath returns the path to the inverted index datastore
func IIndexPath() string {
	return iindexPath
}
