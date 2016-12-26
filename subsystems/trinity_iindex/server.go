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

func startIIndexServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	trinity.RegisterIIndexServer(grpcServer, newIIndexServer())
	grpcServer.Serve(lis)
}

type iindexServer struct {
}

func newIIndexServer() *iindexServer {
	s := new(iindexServer)
	return s
}

func (iis *iindexServer) GetIMap(ctx context.Context, req *trinity.IMapRequest) (*trinity.InvertedMap, error) {
	resultMap, err := iindex.store.GetIMap(req.Termkey)
	return &resultMap, err
}

func (iis *iindexServer) SetIMap(ctx context.Context, req *trinity.InvertedMap) (*trinity.SetResult, error) {
	err := iindex.store.PutIMap(req.Key, req)
	if err != nil {
		return &trinity.SetResult{Success: false}, err
	}
	return &trinity.SetResult{Success: true}, err
}
