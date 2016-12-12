package main

import (
	"net/http"
	"sync"

	"google.golang.org/grpc"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/config"
)

//var gatewayServer *GatewayServer
var gatewayState *GatewayState

var conn *grpc.ClientConn
var gatewayClient trinity.GatewayClient

func RunServer(staticPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	fs := http.FileServer(http.Dir(staticPath))
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/search", searchHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

func main() {

	var mainWaitGroup sync.WaitGroup
	staticPath := config.StaticDir()
	var err error
	gatewayState, err = NewGatewayState(staticPath)
	if err != nil {
		panic(err)
	}

	mainWaitGroup.Add(1)
	go RunServer(staticPath, &mainWaitGroup)

	setupGRPCConn(conn, &gatewayClient)

	mainWaitGroup.Wait()
	defer conn.Close()

}
