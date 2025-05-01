package router

import (
	"net/http"

	"github.com/AnthonyGOliveira/mcp/internal/api/handlers"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	return mux
}
