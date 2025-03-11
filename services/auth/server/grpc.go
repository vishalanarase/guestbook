package server

import (
	"log"
	"net"

	"github.com/vishalanarase/guestbook/clients/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer() {
	// Listen on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the auth service with the gRPC server
	auth.RegisterAuthServiceServer(s, &server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Serve the gRPC server
	log.Print("Auth service started on port :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
