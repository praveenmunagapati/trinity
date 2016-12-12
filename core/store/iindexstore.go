package store

import "github.com/clownpriest/trinity/api/trinity"

/*
InvertIndexStore is a database interface to some underlying inverted
index datastore
*/
type InvertIndexStore interface {
	Put(key, value []byte) error
	PutInvertedMap(key []byte, value *trinity.InvertedMap)
	Get(key []byte) error
	GetInvertedMap(key []byte) *trinity.InvertedMap
}
