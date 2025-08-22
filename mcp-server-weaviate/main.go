package main

import (
	"log"
	"os"
)

func main() {
	// Get port from environment (Railway sets PORT automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8675" // Default port
	}

	s, err := NewMCPServer()
	if err != nil {
		log.Fatalf("failed to start mcp server: %v", err)
	}
	
	// Start HTTP server for Railway deployment
	s.ServeHTTP(port)
}
