package main

import "github.com/clownpriest/trinity/core/config"

/*
IIndex is the global state struct for the inverted
index subsystem
*/
type IIndex struct {
	store  *IIndexStore
	config config.IIndexConfig
}

/*
NewIIndex returns a new forward index object
*/
func NewIIndex(store *IIndexStore, config config.IIndexConfig) IIndex {
	return IIndex{store: store,
		config: config,
	}
}
