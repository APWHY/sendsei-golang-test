package main

import (
	"log"
	"net/http"

	"github.com/burstsms/golang-test/api"
	"github.com/burstsms/golang-test/rpc"
	"github.com/burstsms/golang-test/service"
)

func main() {
	rpcServer, err := rpc.NewServer(service.NewService("API_KEY"), service.Port)
	if err != nil {
		log.Fatalf("failed to initialise service: %s reason: %s\n", service.Name, err)
	}

	go rpcServer.Listen()

	http.HandleFunc("/", api.HelloHandler)
	http.ListenAndServe(":11701", nil)
}
