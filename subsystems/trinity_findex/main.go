package main

import (
	"fmt"
	"os"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/commands"
	"github.com/clownpriest/trinity/core/config"
)

var findex FIndex

func main() {
	if len(os.Args) != 2 {
		fmt.Println(commands.FIndexHelp)
		os.Exit(1)
	}

	config, err := config.LoadFIndexConfig(os.Args[1])
	if err != nil {
		panic(err)
	}
	var store *FIndexStore
	store, err = NewFIndexStore(config)
	if err != nil {
		panic(err)
	}

	findex = NewFIndex(store, config)

	fmap1 := trinity.NewForwardMap()

	fmap1.Put("apple", 7)
	fmap1.Put("apple", 45)
	fmap1.Put("apple", 137)
	fmap1.Put("horse", 55)
	fmap1.Put("clown", 2)
	fmap1.Put("bus", 99)
	fmap1.Put("shoes", 99)

	findex.store.PutFMap("key1", &fmap1)

	getResult, getErr := findex.store.GetFMap("key1")
	if getErr != nil {
		panic(getErr)
	}

	fmt.Println(getResult)

	fmap2 := trinity.NewForwardMap()

	fmap2.Put("bees", 7)
	fmap2.Put("trees", 45)
	fmap2.Put("trees", 137)
	fmap2.Put("blueberry", 55)
	fmap2.Put("clown", 2)
	fmap2.Put("brown", 99)
	fmap2.Put("juniper", 93)
	fmap2.Put("callamity", 90)
	fmap2.Put("special", 99)
	fmap2.Put("humble", 19)

	findex.store.PutFMap("key2", &fmap2)

	getResult2, getErr2 := findex.store.GetFMap("key2")
	if getErr2 != nil {
		panic(getErr2)
	}
	fmt.Println(getResult2)
}
