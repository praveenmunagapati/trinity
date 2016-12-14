package main

import (
	"fmt"
	"os"
	"sync"

	"google.golang.org/grpc"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/commands"
	"github.com/clownpriest/trinity/core/config"
)

var gatewayState *GatewayState

var conn *grpc.ClientConn
var gatewayClient trinity.GatewayClient

var mainConfig config.GatewayConfig

func main() {

	if len(os.Args) < 2 {
		fmt.Println(commands.GatewayHelp)
		os.Exit(0)
	}

	var err error

	mainConfig, err = config.LoadGatewayConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	var mainWaitGroup sync.WaitGroup
	staticPath := config.StaticDir()
	gatewayState, err = NewGatewayState(staticPath)
	if err != nil {
		panic(err)
	}

	mainWaitGroup.Add(1)

	go RunGatewayServer(staticPath, &mainWaitGroup)
	setupGRPCConn(conn, &gatewayClient, mainConfig.ClientPort)

	mainWaitGroup.Wait()
	defer conn.Close()

}
