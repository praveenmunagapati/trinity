package core

import (
	"bytes"
	"os/exec"

	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/system"
)

/*
A TrinityNode struct maintains the global state of a single engine instance
*/
type TrinityNode struct {
	Config         *config.Config
	Subsystems     system.Subsystems
	GatewayCommand *exec.Cmd
}

/*
InitSubsystems launches all the subsystems subprocesses that are specified
in the node's registered subsystems list.
*/
func (node *TrinityNode) InitSubsystems() error {
	for _, system := range node.Subsystems {
		switch system.Role {
		case "gateway":
			gatewayCmd, gateErr := LoadGateway(node.Config)
			if gateErr != nil {
				return gateErr
			}
			node.Subsystems["gateway"].Process = gatewayCmd
		}
	}
	return nil
}

func LoadGateway(config *config.Config) (*exec.Cmd, error) {
	cmd := exec.Command(config.Gateway.BinName, config.Gateway.ConfigPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		return cmd, err
	}
	return cmd, nil
}
