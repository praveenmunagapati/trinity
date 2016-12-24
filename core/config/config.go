package config

import (
	"encoding/json"
	"errors"
	"os"
)

var (

	// ErrConfigDoesntExist means that a supplied path to a
	// config file failed to find anything
	ErrConfigDoesntExist = errors.New("config: config file doesn't exist")
)

const (
	DefaultCoreServerPort   = 50051
	DefaultFIndexServerPort = 50052
	DefaultIIndexServerPort = 50053
)

/*
NodeRole describes the function of the node. trinity
nodes can be standalone full nodes, or be dedicated
to a single function, and serve a wider swarm of nodes.
*/
type NodeRole string

const (
	// Full node performs all engine operations by itself
	Full NodeRole = "full"

	// Crawler node handles crawling the network and queing data for ingestion
	Crawler = "crawler"

	// FIndexer node handles building the forward index
	FIndexer = "findexer"

	// IIndexer node handles building the inverted index
	IIndexer = "iindexer"

	// Gateway node serves as an enduser endpoint, exposing an API
	Gateway = "gateway"
)

/*
Config stores a node's global configuration
parameters. Values can be updated either by
hand in the config file or through the API.
*/
type Config struct {
	// NodeRole represents the function that this node will serve
	NodeRole NodeRole `json:"node_role"`

	// GatewayConfig is the config struct for the gateway server subsystem
	Gateway GatewayConfig `json:"gateway_config"`

	// FIndex is the config stuct for the forward index subsystem
	FIndex FIndexConfig `json:"findex_config"`

	// IIndex is the config struct for the inverted index subsystem
	IIndex IIndexConfig `json:"iindex_config"`

	// APIPort is the port used for the gateway's gRPC
	CoreServerPort int `json:"core_server_port"`

	// SubscribedChannels is a list of data channels that this node
	// is subscribed to
	SubscribedChannels []string `json:"subscribed_channels"`
}

/*
NewDefaultConfig returns a default initialized Config struct.
This is set upon a node's initialization, writing a config file
in the .trinity directory. After the first init, the config will
be loaded from the file.
*/
func NewDefaultConfig() Config {
	return Config{
		NodeRole:       Gateway,
		CoreServerPort: DefaultCoreServerPort,
		SubscribedChannels: []string{
			"YwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			"YwAPJzv6CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			"YwAPJzv7CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
		},
	}
}

/*
LoadRootConfig returns a Config struct loaded from a specified
path.
*/
func LoadRootConfig(path string) (Config, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return Config{}, ErrConfigDoesntExist
	}
	if err == nil {
		configFile, openErr := os.Open(path)

		if openErr != nil {
			return Config{}, err
		}
		decoder := json.NewDecoder(configFile)
		config := Config{}
		decoder.Decode(&config)
		return config, nil
	}

	return Config{}, err
}

/*
GatewayConfig is the configuration struct for the gateway
server subsystem
*/
type GatewayConfig struct {
	BinName    string `json:"bin_name"`
	ClientPort int    `json:"client_port"`
	ConfigPath string `json:"config_path"`
}

/*
NewDefaultGatewayConfig returns an FIndexConfig struct with
values set to their default
*/
func NewDefaultGatewayConfig() GatewayConfig {
	return GatewayConfig{
		BinName:    "trinity_gateway",
		ClientPort: DefaultCoreServerPort,
		ConfigPath: gatewayConfigPath,
	}
}

/*
LoadGatewayConfig returns a configuration struct for the
gateway server subsystem
*/
func LoadGatewayConfig(path string) (GatewayConfig, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return GatewayConfig{}, ErrConfigDoesntExist
	}
	if err == nil {
		configFile, openErr := os.Open(path)
		if openErr != nil {
			return GatewayConfig{}, err
		}
		decoder := json.NewDecoder(configFile)
		config := GatewayConfig{}
		decoder.Decode(&config)
		return config, nil
	}
	return GatewayConfig{}, err
}

/*
FIndexConfig is the configuration struct for the forward
index subsystem
*/
type FIndexConfig struct {
	MaxIndexSize uint   `json:"max_findex_size"`
	StorePath    string `json:"store_path"`
	BinName      string `json:"bin_name"`
	ClientPort   int    `json:"client_port"`
	ConfigPath   string `json:"config_path"`
}

/*
NewDefaultFIndexConfig returns an FIndexConfig struct with
values set to their default
*/
func NewDefaultFIndexConfig() FIndexConfig {
	return FIndexConfig{
		MaxIndexSize: 1000000000,
		BinName:      "trinity_findex",
		ClientPort:   DefaultCoreServerPort,
		ConfigPath:   tfindexConfigPath,
	}
}

/*
LoadFIndexConfig returns a configuration struct for the
forward index subsystem
*/
func LoadFIndexConfig(path string) (FIndexConfig, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return FIndexConfig{}, ErrConfigDoesntExist
	}
	if err == nil {
		configFile, openErr := os.Open(path)
		if openErr != nil {
			return FIndexConfig{}, err
		}
		decoder := json.NewDecoder(configFile)
		config := FIndexConfig{}
		decoder.Decode(&config)
		return config, nil
	}
	return FIndexConfig{}, err
}

/*
IIndexConfig is the configuration struct for the forward
index subsystem
*/
type IIndexConfig struct {
	MaxIndexSize uint   `json:"max_iindex_size"`
	StorePath    string `json:"store_path"`
	BinName      string `json:"bin_name"`
	ClientPort   int    `json:"client_port"`
	ConfigPath   string `json:"config_path"`
}

/*
NewDefaultIIndexConfig returns a configuration struct for
the inverted index subsystem
*/
func NewDefaultIIndexConfig() IIndexConfig {
	return IIndexConfig{
		MaxIndexSize: 1000000000,
		BinName:      "trinity_iindex",
		ClientPort:   DefaultCoreServerPort,
		ConfigPath:   tiindexConfigPath,
	}
}

/*
LoadIIndexConfig returns a configuration struct for
the inverted index subsystem
*/
func LoadIIndexConfig(path string) (IIndexConfig, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return IIndexConfig{}, ErrConfigDoesntExist
	}
	if err == nil {
		configFile, openErr := os.Open(path)
		if openErr != nil {
			return IIndexConfig{}, err
		}
		decoder := json.NewDecoder(configFile)
		config := IIndexConfig{}
		decoder.Decode(&config)
		return config, nil
	}
	return IIndexConfig{}, err
}
