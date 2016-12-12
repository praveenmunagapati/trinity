package api

import (
	"golang.org/x/net/context"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/store/boltstore"
)

/*
FIndexServer is the server struct for the forward index
datastore
*/
type FIndexServer struct {
	store boltstore.BoltStore
	port  string
}

/*
NewFindexServer returns a new forward index server, with
a datastore supplied as argument
*/
func NewFindexServer(store boltstore.BoltStore, port string) FIndexServer {
	return FIndexServer{store: store, port: port}
}

/*
GetDocMap returns a forward index for a document as a proto message
*/
func (fiServe *FIndexServer) GetDocMap(context context.Context, req *trinity.DocMapRequest) (*trinity.ForwardMap, error) {
	key := []byte(req.Docname)
	result, err := fiServe.store.Get(key)
	if err != nil {
		return &trinity.ForwardMap{}, err
	}
	return result, nil
}

func (fiServe *FIndexServer) Store() *boltstore.BoltStore {
	return &fiServe.store
}
