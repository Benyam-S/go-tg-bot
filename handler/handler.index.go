package handler

import (
	"github.com/Benyam-S/go-tg-bot/log"
)

// TelegramBotHandler is a struct that defines a telegram bot handler
type TelegramBotHandler struct {
	debug  bool
	botLog *log.LogBug
}

// BotResponse is a type that defines a bot response message
type BotResponse struct {
	Ok        bool  `json:"ok"`
	ErrorCode int64 `json:"error_code"`
}

// NewTelegramBotHandler is a function that returns a new telegram bot handler
func NewTelegramBotHandler(debugMode bool, botLog *log.LogBug) *TelegramBotHandler {
	if debugMode && botLog == nil {
		botLog = &log.LogBug{
			Logger: &log.Logger{},
		}
	}
	return &TelegramBotHandler{debug: debugMode, botLog: botLog}
}
