package main

type FIndex struct {
	store *FIndexStore
}

/*
NewFIndex returns a new forward index object
*/
func NewFIndex(store *FIndexStore) FIndex {
	return FIndex{store: store}
}
