package store

/*
Store is a database interface to some underlying forward
index datastore
*/
type Store interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	//Contains(key []byte) bool
}
