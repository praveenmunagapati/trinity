package main

import (
	"fmt"
	"net/http"
	"time"
)

func NewGatewayServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-doc/", getDocHandler)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.Handler = mux
	return s
}

func StartMainLoop() {

}

func getDocHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello there")
}
