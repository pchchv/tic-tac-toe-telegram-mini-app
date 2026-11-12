package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tic-tac-toe-miniapp/internal/config"
	"time"
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

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}
