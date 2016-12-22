package main

import (
	"fmt"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/clownpriest/trinity/api/trinity"
)

func startFIndexServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	trinity.RegisterFIndexServer(grpcServer, newFIndexServer())
	grpcServer.Serve(lis)
}

type findexServer struct {
}

func newFIndexServer() *findexServer {
	s := new(findexServer)
	return s
}

func (fis *findexServer) GetDocMap(ctx context.Context, req *trinity.DocMapRequest) (*trinity.ForwardMap, error) {
	resultMap, err := findex.store.GetFMap(req.Docname)
	return &resultMap, err
}

func (fis *findexServer) SetDocMap(ctx context.Context, req *trinity.ForwardMap) (*trinity.SetResult, error) {
	err := findex.store.PutFMap(req.Key, req)
	if err != nil {
		return &trinity.SetResult{Success: false}, err
	}
	return &trinity.SetResult{Success: true}, err
}
