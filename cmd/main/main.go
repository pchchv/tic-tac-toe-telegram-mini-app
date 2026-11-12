package main

import (
	"log"
	"tic-tac-toe-miniapp/internal/config"
)

func main() {
	// Load application configuration
	cfg := config.Load()

	if cfg.BotToken == "" {
		log.Println("Warning: TG_BOT_TOKEN is not set")
	}
}
