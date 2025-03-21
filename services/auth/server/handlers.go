package server

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vishalanarase/guestbook/clients/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	users = map[string]string{} // Map to store users (username -> hashed password)
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// loadRSAPrivateKey function is removed from the code
	privKey, err := loadRSAPrivateKey("../../../keys/auth-private-key.pem")
	if err != nil {
		return nil, fmt.Errorf("could not load private key: %v", err)
	}

	// Sign the token with our secret key
	tokenString, err := token.SignedString(privKey)
	if err != nil {
		return nil, fmt.Errorf("could not create token: %v", err)
	}

	// Return the response
	return &auth.LoginResponse{Token: tokenString, Success: true}, nil
}

// ValidateToken function to validate a JWT token
func (s *server) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {

	// loadRSAPublicKey function is removed from the code
	pubKey, err := loadRSAPublicKey("../../../keys/auth-public-key.pem")
	if err != nil {
		return nil, fmt.Errorf("could not load public key: %v", err)
	}

	// Parse the token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.GetToken(), claims, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	// Return the response
	return &auth.ValidateTokenResponse{Valid: true, Username: claims.Username}, nil
}
