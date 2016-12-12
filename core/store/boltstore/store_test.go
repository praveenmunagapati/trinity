package boltstore

import (
	"fmt"
	"path"
	"strconv"
	"testing"

	"github.com/clownpriest/trinity/core/intake"
	"github.com/clownpriest/trinity/core/intake/htmlintake"
	"github.com/clownpriest/trinity/core/token"
)

func TestBoltStore(t *testing.T) {
	db, _ := NewBoltStore("test_store/index_proto.db")
	defer db.store.Close()

	buffer := &intake.BufferPool{}
	queue := token.NewQueue()

	files, _ := intake.GatherFiletype("../../intake/html/single_testdata", ".html")

	fmt.Printf("\n#files to parse: %d\n\n", len(*files))
	htmlintake.ReadLoop(files, buffer, queue)
	for i, list := range queue.Data() {

		tokenMap := token.NewForwardMap()
		tokenMap.AddList(list)
		fmt.Println(list.Filename())
		db.PutForwardMap(list.Filename(), tokenMap)
		retrievedMap, _ := db.GetForwardMap(list.Filename())
		// brain, _ := retrievedMap.Get("BRAIN")
		// for _, value := range brain.Locations {
		// 	fmt.Println(value)
		// }
		//fmt.Println(retrievedMap)
		retrievedMap.SaveProtoToFIle(path.Join("test_out", strconv.Itoa(i)+".protoout"))
		retrievedMap.SaveJSON(path.Join("test_out", strconv.Itoa(i)+".json"))
		//tokenMap.SaveJSON(path.Join("test_out", strconv.Itoa(i)))
	}
}
