package main

import (
	"log"
	"net"

	"github.com/vishalanarase/guestbook/clients/auth"
	"google.golang.org/grpc"
)

var (
	secretKey = []byte("secretkey") // Secret key for JWT signing
	users     = map[string]string{} // Map to store users (username -> hashed password)
)

type server struct {
	auth.UnimplementedAuthServiceServer
}

func main() {
	// Listen on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the auth service with the gRPC server
	auth.RegisterAuthServiceServer(s, &server{})

	// Serve the gRPC server
	log.Print("Starting server on port :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
