package main

import (
	"sync"

	"google.golang.org/grpc"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/config"
)

var gatewayState *GatewayState

var conn *grpc.ClientConn
var gatewayClient trinity.GatewayClient

func main() {

	var mainWaitGroup sync.WaitGroup
	staticPath := config.StaticDir()
	var err error
	gatewayState, err = NewGatewayState(staticPath)
	if err != nil {
		panic(err)
	}

	mainWaitGroup.Add(1)
	go RunGatewayServer(staticPath, &mainWaitGroup)

	setupGRPCConn(conn, &gatewayClient)

	mainWaitGroup.Wait()
	defer conn.Close()

}
