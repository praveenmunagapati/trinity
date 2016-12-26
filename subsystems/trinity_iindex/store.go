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
	ErrDataNotByteSlice = errors.New("IIndex store: data not a byte slice")
)

type IIndexStore struct {
	rootPath         string
	decompressBuffer []byte
	store            ds.Datastore
}

func NewIIndexStore(config config.IIndexConfig) (*IIndexStore, error) {
	decompBuffer := make([]byte, 1000, 1000)

	fs, err := fs.NewDatastore(config.StorePath)
	if err != nil {
		return &IIndexStore{}, err
	}

	return &IIndexStore{
		rootPath:         config.StorePath,
		store:            fs,
		decompressBuffer: decompBuffer,
	}, nil
}

/*
PutIMap puts a InvertedMap proto struct into the inverted index
data store, assigning it to the term ID key. The InvertedMap
data is compressed before being serialized to disk.
*/
func (iis *IIndexStore) PutIMap(key string, imap *trinity.InvertedMap) error {
	data, err := imap.Marshal()
	if err != nil {
		return err
	}

	var compressed []byte
	compressed, err = zstd.Compress(compressed, data)
	if err != nil {
		return err
	}

	err = iis.store.Put(ds.NewKey(key), compressed)
	if err != nil {
		return err
	}
	return nil
}

/*
GetIMap returns an InvertedMap proto struct from the inverted
index store, given a term lookup key. This function
decompresses the data that's pulled from the data store and
serializes it into a trinity.InvertedMap struct.
*/
func (iis *IIndexStore) GetIMap(key string) (trinity.InvertedMap, error) {
	result, err := iis.store.Get(ds.NewKey(key))
	if err != nil {
		return trinity.InvertedMap{}, err
	}
	var decompressed []byte
	data, assertErr := result.([]byte)
	if !assertErr {
		return trinity.InvertedMap{}, ErrDataNotByteSlice
	}
	decompressed, err = zstd.Decompress(iis.decompressBuffer, data)
	if err != nil {
		return trinity.InvertedMap{}, err
	}
	imapResult := trinity.InvertedMap{}
	imapResult.Unmarshal(decompressed)
	return imapResult, nil
}
