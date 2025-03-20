package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vishalanarase/guestbook/clients/auth"
	"github.com/vishalanarase/guestbook/clients/guestbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Auth struct {
	client auth.AuthServiceClient
}

type server struct {
	guestbook.UnimplementedGuestbookServiceServer
	auth     Auth
	messages map[string][]string
}

// AddMessage
func (s *server) AddMessage(ctx context.Context, req *guestbook.AddMessageRequest) (*guestbook.AddMessageResponse, error) {
	// Validate the token
	resp, err := s.auth.client.ValidateToken(ctx, &auth.ValidateTokenRequest{Token: req.GetToken()})
	if err != nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	// Check if the messages map is nil if so create a new map
	if s.messages == nil {
		s.messages = make(map[string][]string)
	}

	// Check if the user has any messages in the map if not create a new list
	if _, ok := s.messages[resp.GetUsername()]; !ok {
		s.messages[resp.GetUsername()] = []string{}
	}

	// Add the message to the list
	s.messages[resp.GetUsername()] = append(s.messages[resp.GetUsername()], req.GetMessage())

	return &guestbook.AddMessageResponse{Success: true}, nil
}

// GetMessage
func (s *server) GetMessage(ctx context.Context, req *guestbook.GetMessagesRequest) (*guestbook.GetMessageResponse, error) {
	// Validate the token
	resp, err := s.auth.client.ValidateToken(ctx, &auth.ValidateTokenRequest{Token: req.GetToken()})
	if err != nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	// Return the messages
	return &guestbook.GetMessageResponse{Messages: s.messages[resp.Username]}, nil
}

func main() {
	// Create a gRPC connection to the auth service
	authConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial auth server: %v", err)
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
