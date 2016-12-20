package core

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/system"
)

/*
A TrinityNode struct maintains the global state of a single engine instance
*/
type TrinityNode struct {
	Config     *config.Config
	Subsystems []*system.Subsystem
}

/*
InitSubsystems launches all the subsystems subprocesses that are specified
in the node's registered subsystems list.
*/
func (node *TrinityNode) InitSubsystems() error {
	var err error
	for _, system := range node.Subsystems {
		switch system.Role {
		case "gateway":
			err = LoadGateway(node.Config)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func LoadGateway(config *config.Config) error {
	fmt.Println(config.Gateway.BinName)
	fmt.Println(config.Gateway.ConfigPath)
	cmd := exec.Command(config.Gateway.BinName, config.Gateway.ConfigPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Println(out.String())
	if err != nil {
		return err
	}
	return nil
}
