package main

import (
	"log"
	"net/http"
	"os"

	"github.com/burstsms/golang-test/api"
	"github.com/burstsms/golang-test/rpc"
	"github.com/burstsms/golang-test/service"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	rpcServer, err := rpc.NewServer(service.NewService(apiKey), service.Port)
	if err != nil {
		log.Fatalf("failed to initialise service: %s reason: %s\n", service.Name, err)
	}

	go rpcServer.Listen()
	http.HandleFunc("/sms", api.SendHandler)
	http.HandleFunc("/", api.HelloHandler)
	log.Printf("Serving...")
	http.ListenAndServe(":11701", nil)
}
