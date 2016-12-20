package main

import (
	"errors"

	"github.com/clownpriest/trinity/core"
	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/system"
)

var (
	ErrNodeRoleUndefined = errors.New("node role is undefined")
)

/*
LoadNode handles the startup initialization of the
TrinityNode struct. This stuct holds the global state
of the engine node.
*/
func LoadNode(config *config.Config) (*core.TrinityNode, error) {
	node := &core.TrinityNode{}
	node.Config = config
	systems, err := DetermineSubsystems(config)
	if err != nil {
		return node, err
	}

	node.Subsystems = systems
	return node, nil
}

/*
DetermineSubsystems constructs all the Subsystem
structs given the NodeRole supplied in the main
config
*/
func DetermineSubsystems(config *config.Config) (system.Subsystems, error) {
	var systems system.Subsystems
	switch config.NodeRole {
	case "gateway":
		gateway := system.NewSubsystem(system.Gateway)
		systems.Add(&gateway)
	case "findexer":
		findex := system.NewSubsystem(system.ForwardIndex)
		systems.Add(&findex)
	case "iindexer":
		iindex := system.NewSubsystem(system.InvertedIndex)
		systems.Add(&iindex)
	case "full":
		gateway := system.NewSubsystem(system.Gateway)
		findex := system.NewSubsystem(system.ForwardIndex)
		iindex := system.NewSubsystem(system.InvertedIndex)
		systems.Add(&gateway)
		systems.Add(&findex)
		systems.Add(&iindex)
	default:
		return systems, ErrNodeRoleUndefined
	}

	return systems, nil
}

// func InitSubsystems(systems system.Subsystems, config *config.Config) error {
// 	for _, system := range systems {
// 		switch system.Role {
// 		case "gateway":
// 			LoadGateway(config)
// 		}
// 	}
// 	return nil
// }
//
// func LoadGateway(config *config.Config) error {
// 	cmd := exec.Command(config.Gateway.BinName)
// 	err := cmd.Run()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func LoadSubsystems(node *core.TrinityNode) *core.TrinityNode {
// 	return node
// }
