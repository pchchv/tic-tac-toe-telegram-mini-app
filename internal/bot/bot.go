package bot

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Bot wraps the go-telegram/bot instance and application settings
type Bot struct {
	api       *bot.Bot
	webAppURL string
}

// startHandler responds to /start with a WebApp launch button
func (b *Bot) startHandler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Fallback message if WEBAPP_URL is not configured yet
	if b.webAppURL == "" {
		_, _ = tgBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Welcome! The game URL is currently not configured on the server.",
		})
		return
	}

	// Build inline keyboard with WebApp info
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:   "Play Tic-Tac-Toe ❌⭕",
					WebApp: &models.WebAppInfo{URL: b.webAppURL},
				},
			},
		},
	}

	_, err := tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Welcome to Tic-Tac-Toe! Click the button below to start playing:",
		ReplyMarkup: kb,
	})
	if err != nil {
		log.Printf("Failed to send /start message: %v\n", err)
	}
}

// defaultHandler handles all other unrecognized messages
func (b *Bot) defaultHandler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	_, _ = tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Send /start to launch the Tic-Tac-Toe game.",
	})
}
