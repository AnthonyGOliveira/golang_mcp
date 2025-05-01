package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnthonyGOliveira/mcp/internal/api/router"
)

func main() {
	port := 8000
	router := router.SetupRouter()

	log.Printf("Server is running on port %d\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
