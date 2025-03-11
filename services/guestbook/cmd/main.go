package main

import (
	"log"
	"net"

	"github.com/siderolabs/go-api-signature/api/auth"
	"github.com/vishalanarase/guestbook/clients/guestbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Auth struct {
	client auth.AuthServiceClient
}

type server struct {
	guestbook.UnimplementedGuestbookServiceServer
	auth Auth
}

func main() {
	// Create a gRPC connection to the auth service
	authConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial auth service: %v", err)
	}
	defer authConn.Close()

	// Create a new gRPC client for the auth service
	client := auth.NewAuthServiceClient(authConn)

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the guestbook service with the gRPC server
	guestbook.RegisterGuestbookServiceServer(s, &server{
		auth: Auth{
			client: client,
		},
	})

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Serve the gRPC server
	log.Print("Guestbook service started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
