package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	port := 8000
	log.Printf("Server is running on port %d\n", port)

	http.HandleFunc("/health", healthCheckHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
