package store

import "github.com/clownpriest/trinity/api/trinity"

/*
ForwardIndexStore is an interface to a forward index
store, using some unspecified storage engine.
*/
type ForwardIndexStore interface {
	Put(key, value []byte) error
	PutForwardMap(key string, value *trinity.ForwardMap) error
	Get(key []byte) ([]byte, error)
	GetForwardMap(key []byte) (*trinity.ForwardMap, error)
}
