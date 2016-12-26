package main

import "github.com/clownpriest/trinity/core/config"

/*
FIndex is the global state struct for the ForwardIndex
subsystem
*/
type FIndex struct {
	store  *FIndexStore
	config config.FIndexConfig
}

/*
NewFIndex returns a new forward index object
*/
func NewFIndex(store *FIndexStore, config config.FIndexConfig) FIndex {
	return FIndex{store: store,
		config: config,
	}
}
