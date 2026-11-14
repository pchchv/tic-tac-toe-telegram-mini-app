package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tic-tac-toe-miniapp/internal/bot"
	"tic-tac-toe-miniapp/internal/config"
	"time"
)

func main() {
	// Load application configuration
	cfg := config.Load()

	// Setup root context with cancellation for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize and start Telegram bot if token is provided
	switch cfg.BotToken {
	case "":
		log.Println("Warning: TG_BOT_TOKEN is not set, skipping bot startup")
	case "your_bot_token_here":
		log.Println("Warning: Valid TG_BOT_TOKEN is not set, skipping bot startup")
	default:
		tgBot, err := bot.New(cfg.BotToken, cfg.WebAppURL)
		if err != nil {
			log.Fatalf("Failed to initialize Telegram bot: %v\n", err)
		}
		go tgBot.Start(ctx)
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

	// Start HTTP server in a goroutine
	go func() {
		log.Printf("Server is starting on %s\n", cfg.ListenAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	// Wait for termination signal
	<-ctx.Done()

	log.Println("Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Application gracefully stopped")
}
