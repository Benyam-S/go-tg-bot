package handler

import (
	"strings"

	"github.com/Benyam-S/go-tg-bot/log"
)

// TelegramBotHandler is a struct that defines a telegram bot handler
type TelegramBotHandler struct {
	BotAPIAccessPoint string
	BotAccessToken    string
	BotID             int64
	BotUsername       string
	BotURL            string
	logger            log.ILogger
	logs              *log.LogContainer // logs can never be nil
}

// NewTelegramBotHandler is a function that returns a new telegram bot handler
func NewTelegramBotHandler(botAPIAccessPoint string, botAccessToken string, botID int64, botUsername string,
	botLogger log.ILogger, botLogs *log.LogContainer) *TelegramBotHandler {

	var botURL string

	// checking if the bot username starts with "@"
	if len(strings.TrimSpace(botUsername)) > 0 {
		if strings.HasPrefix(botUsername, "@") {
			botURL = "https://t.me/" + botUsername[1:]
		} else {
			botUsername = "@" + botUsername
			botURL = "https://t.me/" + botUsername
		}
	}

	return &TelegramBotHandler{BotAPIAccessPoint: botAPIAccessPoint, BotAccessToken: botAccessToken,
		BotID: botID, BotURL: botURL, BotUsername: botUsername, logger: botLogger, logs: botLogs}
}

// Logging is a method that will be internally used for making logging efficient
func (handler *TelegramBotHandler) Logging(stmt, logFile string) {
	if handler.logger != nil {
		if handler.logs != nil {
			if logFile == log.ErrorLogFile {
				logFile = handler.logs.ErrorLogFile
			} else {
				logFile = handler.logs.BotLogFile
			}
		}
		handler.logger.Log(stmt, logFile)
	}
}
