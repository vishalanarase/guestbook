package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vishalanarase/guestbook/clients/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	secretKey = []byte("secretkey") // Secret key for JWT signing
	users     = map[string]string{} // Map to store users (username -> hashed password)
)

type server struct {
	auth.UnimplementedAuthServiceServer
}

func (s *server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// Check if user already exists
	if _, exists := users[req.GetUsername()]; exists {
		return nil, fmt.Errorf("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	// Store the user in the map
	users[req.Username] = string(hashedPassword)

	// Return the response
	return &auth.RegisterResponse{Success: true}, nil
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
	log.Print("Auth service started on port :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
