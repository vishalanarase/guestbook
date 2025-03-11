package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vishalanarase/guestbook/clients/auth"
	"github.com/vishalanarase/guestbook/clients/guestbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Auth struct {
	client auth.AuthServiceClient
}

type server struct {
	guestbook.UnimplementedGuestbookServiceServer
	auth     Auth
	messages []string
}

// AddMessage
func (s *server) AddMessage(ctx context.Context, req *guestbook.AddMessageRequest) (*guestbook.AddMessageResponse, error) {
	// Validate the token
	_, err := s.auth.client.ValidateToken(ctx, &auth.ValidateTokenRequest{Token: req.GetToken()})
	if err != nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	// Add the message to the list
	s.messages = append(s.messages, req.GetMessage())

	return &guestbook.AddMessageResponse{Success: true}, nil
}

// GetMessage
func (s *server) GetMessage(ctx context.Context, req *guestbook.GetMessagesRequest) (*guestbook.GetMessageResponse, error) {
	// Validate the token
	_, err := s.auth.client.ValidateToken(ctx, &auth.ValidateTokenRequest{Token: req.GetToken()})
	if err != nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	// Return the messages
	return &guestbook.GetMessageResponse{Messages: s.messages}, nil
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
