package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vishalanarase/guestbook/clients/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	secretKey = []byte("secretkey") // Secret key for JWT signing
	users     = map[string]string{} // Map to store users (username -> hashed password)
)

// Struct for the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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

// Login function to authenticate a user
func (s *server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// Get the hashed password for the user
	hashedPassword, exists := users[req.Username]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Generate a JWT token
	claims := &Claims{
		Username: req.GetUsername(),
		StandardClaims: jwt.StandardClaims{
			Audience:  req.GetUsername(),
			Issuer:    "guestbook-app",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("could not create token: %v", err)
	}

	// Return the response
	return &auth.LoginResponse{Token: tokenString, Success: true}, nil
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

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Serve the gRPC server
	log.Print("Auth service started on port :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
