package main

import (
	"errors"

	"github.com/DataDog/zstd"
	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/config"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/fs"
)

var (
	ErrDataNotByteSlice = errors.New("findex store: data not a byte slice")
)

type FIndexStore struct {
	rootPath         string
	decompressBuffer []byte
	store            ds.Datastore
}

func NewFIndexStore(config config.FIndexConfig) (*FIndexStore, error) {
	decompBuffer := make([]byte, 1000, 1000)

	fs, err := fs.NewDatastore(config.StorePath)
	if err != nil {
		return &FIndexStore{}, err
	}

	return &FIndexStore{
		rootPath:         config.StorePath,
		store:            fs,
		decompressBuffer: decompBuffer,
	}, nil
}

/*
PutFMap puts a ForwardMap proto struct into the forward index
data store, assigning it to the document ID key. The ForwardMap
data is compressed before being serialized to disk.
*/
func (fis *FIndexStore) PutFMap(key string, fmap *trinity.ForwardMap) error {
	data, err := fmap.Marshal()
	if err != nil {
		return err
	}

	var compressed []byte
	compressed, err = zstd.Compress(compressed, data)
	if err != nil {
		return err
	}

	err = fis.store.Put(ds.NewKey(key), compressed)
	if err != nil {
		return err
	}
	return nil
}

/*
GetFMap returns a ForwardMap proto struct from the forward
index store, given a document lookup key. This function
decompresses the data that's pulled from the data store and
serializes it into a trinity.ForwardMap struct.
*/
func (fis *FIndexStore) GetFMap(key string) (trinity.ForwardMap, error) {
	result, err := fis.store.Get(ds.NewKey(key))
	if err != nil {
		return trinity.ForwardMap{}, err
	}
	var decompressed []byte
	data, assertErr := result.([]byte)
	if !assertErr {
		return trinity.ForwardMap{}, ErrDataNotByteSlice
	}
	decompressed, err = zstd.Decompress(fis.decompressBuffer, data)
	if err != nil {
		return trinity.ForwardMap{}, err
	}
	fmapResult := trinity.ForwardMap{}
	fmapResult.Unmarshal(decompressed)
	return fmapResult, nil
}
