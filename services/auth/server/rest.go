package server

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vishalanarase/guestbook/clients/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// gRPC Gateway HTTP server implementation
func StartRESTServer() {
	// Create a new HTTP mux for the gRPC Gateway
	mux := http.NewServeMux()

	// Custom non-gRPC HTTP handler (e.g., health check)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Create a connection to the gRPC server
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC Gateway mux
	gwmux := runtime.NewServeMux()
	err = auth.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("failed to register gRPC Gateway handler: %v", err)
	}

	// Mount the gRPC Gateway on the HTTP mux
	mux.Handle("/", gwmux)

	// Start the HTTP/REST server
	log.Println("HTTP/REST server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}
}
