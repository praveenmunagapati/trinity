package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core"
	"github.com/clownpriest/trinity/core/commands"
	"github.com/clownpriest/trinity/core/commands/cli"
	"github.com/clownpriest/trinity/core/config"
	"github.com/clownpriest/trinity/core/system"
	"github.com/clownpriest/trinity/core/system/subsystem"
)

// GitHash is the git HEAD commit hash the binary
// was built with
var GitHash string

var interuptHandler *system.InteruptHandler

var node *core.TrinityNode

var mainConfig config.Config

type gatewayServer struct {
}

func (gs *gatewayServer) GetSearchQuery(ctx context.Context, query *trinity.SearchQuery) (*trinity.SearchResponse, error) {
	fmt.Printf("got a search request: %s\n", query.Query)

	response := subsystem.Search(query)

	return response, nil
}

func newGatewayServer() *gatewayServer {
	s := new(gatewayServer)
	// s.loadFeatures(*jsonDBFile)
	// s.routeNotes = make(map[string][]*pb.RouteNote)
	return s
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println(commands.RootHelp)
		os.Exit(0)
	}

	commandResult, commandErr := cli.Parse(os.Args)
	if commandErr != nil {
		printUsage()
		os.Exit(0)
	}

	handleCommand(commandResult)

	var err error
	node, err = LoadNode(&mainConfig)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 12345))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	trinity.RegisterGatewayServer(grpcServer, newGatewayServer())
	grpcServer.Serve(lis)

}
