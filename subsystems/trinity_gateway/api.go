package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/clownpriest/trinity/api/trinity"
	uuid "gx/ipfs/QmcyaFHbyiZfoX5GTpcqqCPYmbjYNAhRDekXSJPFHdYNSV/go.uuid"
)

/*
RunGatewayServer starts the http server for the gateway
subprocess.
*/
func RunGatewayServer(staticPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	fs := http.FileServer(http.Dir(staticPath))
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/search", searchHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

/*
SearchResult is the type carrying the response
from the main trinity node
*/
type SearchResult struct {
	Query   string
	Results []*trinity.SearchResult
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(gatewayState.mainPage)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	searchString := r.URL.Query().Get("q")

	if searchString == "" {
		searchString = " "
	}
	fmt.Println(searchString)

	requestID := uuid.NewV5(uuid.NamespaceDNS, searchString)

	req := trinity.SearchRequest{Query: searchString, Id: requestID.Bytes()}

	requestResults, err := gatewayClient.Search(context.Background(), &req)

	results := SearchResult{
		Query:   searchString,
		Results: requestResults.Results,
	}

	gatewayState.searchTemplate.Execute(w, results)
}
