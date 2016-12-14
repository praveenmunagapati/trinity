package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"path"

	"google.golang.org/grpc"

	"github.com/clownpriest/trinity/api/trinity"
	"github.com/clownpriest/trinity/core/config"
)

type GatewayState struct {
	mainPage       []byte
	searchPage     []byte
	searchTemplate *template.Template
}

/*
NewGatewayState returns a global state object for the gateway
*/
func NewGatewayState(staticPath string) (*GatewayState, error) {
	state := &GatewayState{}
	mainPageData, err := readMainPage()
	if err != nil {
		return state, err
	}

	searchPageData, err := readSearchPage()
	if err != nil {
		return state, err
	}

	searchTemplate, err := readSearchTemplate()
	if err != nil {
		return state, err
	}

	state.mainPage = mainPageData
	state.searchPage = searchPageData
	state.searchTemplate = searchTemplate
	return state, nil
}

func setupGRPCConn(conn *grpc.ClientConn, client *trinity.GatewayClient) {
	var dialErr error
	conn, dialErr = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if dialErr != nil {
		log.Fatalf("did not connect: %v", dialErr)
	}

	gatewayClient = trinity.NewGatewayClient(conn)
}

func readMainPage() ([]byte, error) {
	data, err := ioutil.ReadFile(path.Join(config.StaticDir(), "index.html"))
	return data, err
}

func readSearchPage() ([]byte, error) {
	data, err := ioutil.ReadFile(path.Join(config.StaticDir(), "search.html"))
	return data, err
}

func readSearchTemplate() (*template.Template, error) {
	t, err := template.ParseFiles(path.Join(config.StaticDir(), "search.html"))
	return t, err
}
