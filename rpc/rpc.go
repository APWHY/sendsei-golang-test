package rpc

import (
	"log"
	"net"
	"net/rpc"
	"strconv"
)

// Server is the generic wrapper for a service over rpc
type Server struct {
	port    int
	server  *rpc.Server
	service Service
}

// NewServer handles the general set up of the service
// it creates a db instance, net listener and wires up the rpc server
func NewServer(service Service, port int) (*Server, error) {
	server := Server{
		port:    port,
		server:  rpc.NewServer(),
		service: service,
	}

	err := server.server.RegisterName(service.Name(), service.Receiver())

	return &server, err
}

// Listen sets the rpc server to accept connections
func (s *Server) Listen() {
	port := ":" + strconv.Itoa(s.port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to open port for service: %s reason: %s", s.service.Name(), err)
	}
	defer listener.Close()

	s.server.Accept(listener)
}

// Service types are required to have these
type Service interface {
	Name() string
	Receiver() interface{}
}

// Client is a helper to wrap the rpc calls
type Client struct {
	ServiceAddress string
	ServiceName    string
}

func (c *Client) Call(method string, params interface{}, reply interface{}) error {
	log.Printf("rpc: %s.%s\n", c.ServiceName, method)

	client, err := rpc.Dial("tcp", c.ServiceAddress)
	if err != nil {
		return err
	}

	defer client.Close()

	return client.Call(c.ServiceName+"."+method, params, reply)
}
