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

	// Setup HTTP server
	server := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server is starting on %s\n", cfg.ListenAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()
}
