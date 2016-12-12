package main

import (
	"github.com/DataDog/zstd"
	"github.com/clownpriest/trinity/api/trinity"
	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
)

/*
InvertedBoltStore is a DB with  BoltDB backing store
*/
type InvertedBoltStore struct {
	store      *bolt.DB
	mainBucket []byte
}

/*
NewInvertedBoltStore returns a default initialized
InvertedBoltStore opened from a given path
*/
func NewInvertedBoltStore(path string) (InvertedBoltStore, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return InvertedBoltStore{}, err
	}

	boltstore := InvertedBoltStore{store: db,
		mainBucket: []byte("mainIndex"),
	}

	err = boltstore.store.Update(func(tx *bolt.Tx) error {
		_, bucketErr := tx.CreateBucketIfNotExists(boltstore.mainBucket)
		if bucketErr != nil {
			return bucketErr
		}
		return nil
	})
	return boltstore, err
}

/*
Put assigns a key to a value
*/
func (db *InvertedBoltStore) Put(key, value []byte) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.mainBucket)
		putErr := b.Put(key, value)
		return putErr
	})
	return err
}

/*
Get returns a byte slice value assigned to a given key
*/
func (db *InvertedBoltStore) Get(key []byte) ([]byte, error) {
	var result []byte
	err := db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.mainBucket)
		result = b.Get(key)
		return nil
	})
	return result, err
}

/*
GetIMap returns a ForwardMap given a key to document
*/
func (db *InvertedBoltStore) GetIMap(key []byte) (*trinity.ForwardMap, error) {
	result, err := db.Get([]byte(key))
	if err != nil {
		return &trinity.ForwardMap{}, err
	}
	newMap := &trinity.ForwardMap{}
	decompResult, _ := zstd.Decompress(nil, result)
	proto.Unmarshal(decompResult, newMap)
	return newMap, nil
}

/*
PutIMap assigns a trinity.ForwardMap as value to a filename key
*/
func (db *InvertedBoltStore) PutIMap(filename string, tokenMap trinity.ForwardMap) error {
	//var buffer = new(bytes.Buffer)
	data, err := tokenMap.Marshal()
	if err != nil {
		return err
	}
	compressed, _ := zstd.CompressLevel(nil, data, zstd.BestCompression)

	err = db.Put([]byte(filename), compressed)
	if err != nil {
		return err
	}
	return nil
}
