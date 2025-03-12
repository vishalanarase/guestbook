package server

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
)

// Load RSA private key from PEM file
func loadRSAPrivateKey(filepath string) (*rsa.PrivateKey, error) {
	privKeyData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKeyData)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

// Load RSA public key from PEM file
func loadRSAPublicKey(filepath string) (*rsa.PublicKey, error) {
	pubKeyData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyData)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}
