package boltstore

import (
	"errors"

	"github.com/DataDog/zstd"
	"github.com/boltdb/bolt"
	"github.com/clownpriest/trinity/api/trinity"
	"github.com/gogo/protobuf/proto"
)

var (
	// ErrKeyDoesntExist means that the key provided to Get() returned nil
	ErrKeyDoesntExist = errors.New("key doesn't exist")
)

/*
BoltStore is a DB with  BoltDB backing store
*/
type BoltStore struct {
	store      *bolt.DB
	mainBucket []byte
}

/*
NewBoltStore returns a default initialized BoltStore
opened from a given path
*/
func NewBoltStore(path string) (BoltStore, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return BoltStore{}, err
	}

	boltstore := BoltStore{store: db,
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
func (db *BoltStore) Put(key, value []byte) error {
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
func (db *BoltStore) get(key []byte) ([]byte, error) {
	var result []byte
	err := db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.mainBucket)
		result = b.Get(key)
		return nil
	})
	return result, err
}

func (db *BoltStore) Get(key []byte) (*trinity.ForwardMap, error) {
	result, err := db.get([]byte(key))
	if err != nil {
		return &trinity.ForwardMap{}, err
	}
	newMap := &trinity.ForwardMap{}
	decompResult, _ := zstd.Decompress(nil, result)
	proto.Unmarshal(decompResult, newMap)
	return newMap, nil
}

/*
Contains checks if the Db contains an entry for a key
*/
func (db *BoltStore) Contains(key []byte) bool {
	var exists = true
	db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.mainBucket)
		result := b.Get(key)
		if result == nil {
			exists = false
		}
		return nil
	})
	return exists
}

// /*
// DumpMapJSON assigns a token.Map as value to a filename key
// */
// func (db *BoltStore) DumpMapJSON(filename string, tokenMap token.Map) error {
// 	var buffer = new(bytes.Buffer)
// 	//fmt.Println(tokenMap)
// 	encoder := json.NewEncoder(buffer)
// 	encoder.SetIndent("", "  ")
// 	err := encoder.Encode(tokenMap)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Put([]byte(filename), buffer.Bytes())
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// /*
// DumpMapGob assigns a token.Map as value to a filename key
// */
// func (db *BoltStore) DumpMapGob(filename string, tokenMap token.Map) error {
// 	var buffer = new(bytes.Buffer)
// 	//fmt.Println(tokenMap)
// 	encoder := gob.NewEncoder(buffer)
// 	//encoder.SetIndent("", "  ")
// 	err := encoder.Encode(tokenMap)
// 	if err != nil {
// 		return err
// 	}
//
// 	compressed, _ := zstd.Compress(nil, buffer.Bytes())
// 	err = db.Put([]byte(filename), compressed)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

/*
PutForwardMap assigns a token.ForwardMap as value to a filename key
*/
func (db *BoltStore) PutForwardMap(filename string, tokenMap trinity.ForwardMap) error {
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

/*
GetForwardMap returns a token.ForwardMap that is resolved from a
filename as key
*/
func (db *BoltStore) GetForwardMap(filename string) (*trinity.ForwardMap, error) {
	result, err := db.get([]byte(filename))
	if err != nil {
		return &trinity.ForwardMap{}, err
	}
	newMap := &trinity.ForwardMap{}
	decompResult, _ := zstd.Decompress(nil, result)
	proto.Unmarshal(decompResult, newMap)
	return newMap, nil
}
