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
	err = InitSubsystems(systems)
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
		findex := system.NewSubsystem(system.ForwardIndex)
		iindex := system.NewSubsystem(system.InvertedIndex)
		systems.Add(&findex)
		systems.Add(&iindex)
	default:
		return systems, ErrNodeRoleUndefined
	}

	return systems, nil
}

func InitSubsystems(systems system.Subsystems) error {
	return nil
}

// func LoadSubsystems(node *core.TrinityNode) *core.TrinityNode {
// 	return node
// }
