package bot

import "github.com/go-telegram/bot"

// Bot wraps the go-telegram/bot instance and application settings
type Bot struct {
	api       *bot.Bot
	webAppURL string
}
