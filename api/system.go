package api

import "github.com/clownpriest/trinity/core"

type SystemServer struct {
	system *core.TrinityNode
}

func (sysServe *SystemServer) GetMaxFIndexSize() {
	//	maxFISize := sysServe.system.Config.MaxFIndexSize
}
