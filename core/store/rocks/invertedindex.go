package rocks

import (
	"github.com/DataDog/zstd"
	"github.com/clownpriest/trinity/api/trinity"
	"github.com/gogo/protobuf/proto"
	"github.com/tecbot/gorocksdb"
)

/*
InvertedRocksStore is the storage stuct for a
rocksdb backed inverted index store.
*/
type InvertedStore struct {
	store     *gorocksdb.DB
	writeOpts *gorocksdb.WriteOptions
	readOpts  *gorocksdb.ReadOptions
}

/*
NewInvertedRocksStore return default rocksdb initialized InvertedRocksStore
*/
func NewInvertedStore() (InvertedStore, error) {
	store := InvertedStore{}
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	store.writeOpts = gorocksdb.NewDefaultWriteOptions()
	store.readOpts = gorocksdb.NewDefaultReadOptions()

	db, err := gorocksdb.OpenDb(opts, "iindex_store")
	if err != nil {
		return store, err
	}
	store.store = db
	return store, nil
}

/*
Put key/value into rocks store
*/
func (store *InvertedStore) Put(key, value []byte) error {
	err := store.store.Put(store.writeOpts, key, value)
	return err
}

func (store *InvertedStore) PutInvertedMap(filename string, tokenMap trinity.InvertedMap) error {
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
func (store *InvertedStore) Get(key []byte) ([]byte, error) {
	result, err := store.store.Get(store.readOpts, key)
	if err != nil {
		return []byte{}, err
	}
	return result.Data(), nil
}

/*
GetInvertedMap from the rocks store
*/
func (store *InvertedStore) GetInvertedMap(key []byte) (*trinity.InvertedMap, error) {
	result, err := store.Get([]byte(key))
	if err != nil {
		return &trinity.InvertedMap{}, err
	}
	newMap := &trinity.InvertedMap{}
	decompResult, _ := zstd.Decompress(nil, result)
	proto.Unmarshal(decompResult, newMap)
	return newMap, nil
}
