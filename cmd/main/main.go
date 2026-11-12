package main

import (
	"log"
	"net/http"
	"tic-tac-toe-miniapp/internal/config"
)

func main() {
	// Load application configuration
	cfg := config.Load()

	if cfg.BotToken == "" {
		log.Println("Warning: TG_BOT_TOKEN is not set")
	}

	// Initialize HTTP router
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok", "message": "Backend is running"}`))
	})
}
