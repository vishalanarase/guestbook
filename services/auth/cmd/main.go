package main

import (
	"github.com/vishalanarase/guestbook/services/auth/server"
)

func main() {
	// Start the gRPC servers
	go server.StartGRPCServer()

	// Start the HTTP/REST server
	server.StartRESTServer()
}
