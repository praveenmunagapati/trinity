package rocks

import (
	"github.com/DataDog/zstd"
	"github.com/clownpriest/trinity/api/trinity"
	"github.com/gogo/protobuf/proto"
	"github.com/tecbot/gorocksdb"
)

/*
ForwardRocksStore is the storage stuct for a
rocksdb backed inverted index store.
*/
type ForwardStore struct {
	store     *gorocksdb.DB
	writeOpts *gorocksdb.WriteOptions
	readOpts  *gorocksdb.ReadOptions
}

/*
NewForwardRocksStore return default rocksdb initialized InvertedRocksStore
*/
func NewForwardStore() (ForwardStore, error) {
	store := ForwardStore{}
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	store.writeOpts = gorocksdb.NewDefaultWriteOptions()
	store.readOpts = gorocksdb.NewDefaultReadOptions()

	db, err := gorocksdb.OpenDb(opts, "findex_store")
	if err != nil {
		return store, err
	}
	store.store = db
	return store, nil
}

/*
Put key/value into rocks store
*/
func (store *ForwardStore) Put(key, value []byte) error {
	err := store.store.Put(store.writeOpts, key, value)
	return err
}

/*
PutForwardMap inserts a trinity.ForwardMap into the rocks store.
The function zstd compresses the bytes before serialization.
*/
func (store *ForwardStore) PutForwardMap(filename string, tokenMap *trinity.ForwardMap) error {
	data, err := tokenMap.Marshal()
	if err != nil {
		return err
	}
	compressed, _ := zstd.CompressLevel(nil, data, zstd.BestCompression)

	err = store.Put([]byte(filename), compressed)
	if err != nil {
		return err
	}
	return nil
}

/*
Get key from rocks store
*/
func (store *ForwardStore) Get(key []byte) ([]byte, error) {
	result, err := store.store.Get(store.readOpts, key)
	if err != nil {
		return []byte{}, err
	}
	return result.Data(), nil
}

/*
GetForwardMap from the rocks store. Function handles zstd decompression and
proto unmarshaling
*/
func (store *ForwardStore) GetForwardMap(key []byte) (*trinity.ForwardMap, error) {
	result, err := store.Get([]byte(key))
	if err != nil {
		return &trinity.ForwardMap{}, err
	}
	newMap := &trinity.ForwardMap{}
	decompResult, _ := zstd.Decompress(nil, result)
	proto.Unmarshal(decompResult, newMap)
	return newMap, nil
}
