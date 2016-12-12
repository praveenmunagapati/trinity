package api

import (
	"fmt"

	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/store/boltstore"
)

type ServerGroup struct {
	FIndexServer FIndexServer
	//	iindexServer IIndexServer
}

/*
StartDaemon sets up the node's processes and prepares it to
accept inputs through API
*/
func StartDaemon(mainConfig config.Config) (ServerGroup, error) {
	fmt.Println(mainConfig)
	findexStore, err := boltstore.NewBoltStore(mainConfig.FIndex.StorePath)
	if err != nil {
		return ServerGroup{}, err
	}
	//	iindexStore, _ := boltstore.NewBoltStore(config.IIndexPath)
	serveGroup := ServerGroup{FIndexServer: NewFindexServer(findexStore, mainConfig.APIPort)}
	return serveGroup, nil
}
