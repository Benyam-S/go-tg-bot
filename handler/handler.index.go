package handler

import "github.com/Benyam-S/go-tg-bot/log"

// TelegramBotHandler is a struct that defines a telegram bot handler
type TelegramBotHandler struct {
	logger *log.Logger
}

// BotResponse is a type that defines a bot response message
type BotResponse struct {
	Ok        bool  `json:"ok"`
	ErrorCode int64 `json:"error_code"`
}

// NewTelegramBotHandler is a function that returns a new telegram bot handler
func NewTelegramBotHandler(log *log.Logger) *TelegramBotHandler {
	return &TelegramBotHandler{logger: log}
}
