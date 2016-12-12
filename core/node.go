package core

import (
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
