package main

import (
	"fmt"
	"net"
	"sync"

	"golang.org/x/net/context"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/system/subsystem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func startTrinityCoreServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	trinity.RegisterGatewayServer(grpcServer, newGatewayServer())
	grpcServer.Serve(lis)
}

type gatewayServer struct {
}

func (gs *gatewayServer) GetSearchQuery(ctx context.Context, query *trinity.SearchQuery) (*trinity.SearchResponse, error) {
	fmt.Printf("got a search request: %s\n", query.Query)

	response := subsystem.Search(query)

	return response, nil
}

func newGatewayServer() *gatewayServer {
	s := new(gatewayServer)
	return s
}
