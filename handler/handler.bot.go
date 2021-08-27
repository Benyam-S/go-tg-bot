package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Benyam-S/go-tg-bot/entity"
	"github.com/Benyam-S/go-tg-bot/log"
)

// SendReplyToTelegramChat sends a reply to the Telegram chat identified by its chat Id
// For removing the reply keyboard, use ReplyKeyboardRemove{RemoveKeyboard: true} object as 'reply' string
func (handler *TelegramBotHandler) SendReplyToTelegramChat(chatID int64, reply ...string) (string, error) {

	text := ""
	replyMarkup := ""

	if len(reply) > 0 {
		text = reply[0]
	}

	if len(reply) > 1 {
		replyMarkup = reply[1]
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started sending replay to telegram chat { Chat ID : %d, Text : %s, Markup : %s }",
		chatID, text, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendMessage"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":      {strconv.FormatInt(chatID, 10)},
			"text":         {text},
			"reply_markup": {replyMarkup},
			"parse_mode":   {"html"},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending replay to telegram chat { Chat ID : %d, Text : %s, Markup : %s }, %s",
			chatID, text, replyMarkup, err.Error()), log.ErrorLogFile)

		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending replay to telegram chat { Chat ID : %d, Text : %s, Markup : %s }, %s",
			chatID, text, replyMarkup, errRead.Error()), log.ErrorLogFile)

		return "", errRead
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished sending replay to telegram chat { Chat ID : %d, Text : %s, Markup : %s, Response : %s }",
		chatID, text, replyMarkup, string(bodyBytes)), log.BotLogFile)

	return string(bodyBytes), nil
}

// SendDocumentToTelegramChat sends a document to the Telegram chat identified by its chat Id
func (handler *TelegramBotHandler) SendDocumentToTelegramChat(chatID int64, fileID string, reply ...string) (string, error) {

	caption := ""
	replyMarkup := ""

	if len(reply) > 0 {
		caption = reply[0]
	}

	if len(reply) > 1 {
		replyMarkup = reply[1]
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started sending document to telegram chat { Chat ID : %d, File ID : %s, Caption : %s, Markup : %s }",
		chatID, fileID, caption, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendDocument"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":      {strconv.FormatInt(chatID, 10)},
			"document":     {fileID},
			"caption":      {caption},
			"parse_mode":   {"html"},
			"reply_markup": {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending document to telegram chat { Chat ID : %d, File ID : %s, Caption : %s, Markup : %s }, %s",
			chatID, fileID, caption, replyMarkup, err.Error()), log.ErrorLogFile)

		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending document to telegram chat { Chat ID : %d, File ID : %s, Caption : %s, Markup : %s }, %s",
			chatID, fileID, caption, replyMarkup, errRead.Error()), log.ErrorLogFile)

		return "", err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished sending document to telegram chat { Chat ID : %d, File ID : %s, Caption : %s, Markup : %s, Response : %s }",
		chatID, fileID, caption, replyMarkup, string(bodyBytes)), log.BotLogFile)

	return string(bodyBytes), nil
}

// PostToTelegramChannel posts a certain content to a telegram channel
func (handler *TelegramBotHandler) PostToTelegramChannel(post ...string) (string, error) {

	text := ""
	replyMarkup := ""

	if len(post) > 0 {
		text = post[0]
	}

	if len(post) > 1 {
		replyMarkup = post[1]
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started posting to telegram channel { Channel Name : %s, text : %s, Markup : %s }",
		handler.ChannelName, text, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendMessage"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":      {handler.ChannelName},
			"text":         {text},
			"reply_markup": {replyMarkup},
			"parse_mode":   {"html"},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For posting to telegram channel { Channel Name : %s, text : %s, Markup : %s }, %s",
			handler.ChannelName, text, replyMarkup, err.Error()), log.ErrorLogFile)

		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For posting to telegram channel { Channel Name : %s, text : %s, Markup : %s }, %s",
			handler.ChannelName, text, replyMarkup, errRead.Error()), log.ErrorLogFile)

		return "", err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished posting to telegram channel { Channel Name : %s, text : %s, Markup : %s, Response : %s }",
		handler.ChannelName, text, replyMarkup, string(bodyBytes)), log.BotLogFile)

	return string(bodyBytes), nil
}

// AnswerToTelegramCallBack sends a reply to the Telegram call back request identified by the query id
func (handler *TelegramBotHandler) AnswerToTelegramCallBack(queryID string, text string) (string, error) {

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started answering to telegram callback { Query ID : %s, text : %s }",
		queryID, text), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/answerCallbackQuery"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"callback_query_id": {queryID},
			"text":              {text},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For answering to telegram callback { Query ID : %s, text : %s }, %s",
			queryID, text, err.Error()), log.ErrorLogFile)

		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For answering to telegram callback { Query ID : %s, text : %s }, %s",
			queryID, text, errRead.Error()), log.ErrorLogFile)

		return "", err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished answering to telegram callback { Query ID : %s, text : %s, Response : %s }",
		queryID, text, string(bodyBytes)), log.BotLogFile)

	return string(bodyBytes), nil
}

// CreateReplyKeyboard is a function that creates a reply keyboard from set of parameters
func (handler *TelegramBotHandler) CreateReplyKeyboard(resizeKeyboard, oneTimeKeyboard bool, keyboardButtons ...[]string) string {

	buttonRows := make([][]*entity.ReplyKeyboardButton, 0)

	for _, keyboardRow := range keyboardButtons {

		row := make([]*entity.ReplyKeyboardButton, 0)
		for _, keyboardText := range keyboardRow {
			button := new(entity.ReplyKeyboardButton)
			button.Text = keyboardText
			row = append(row, button)
		}

		buttonRows = append(buttonRows, row)
	}

	keyboard := entity.ReplyKeyboardMarkup{
		Keyboard:        buttonRows,
		ResizeKeyboard:  resizeKeyboard,
		OneTimeKeyboard: oneTimeKeyboard,
	}

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}

// CreateReplyKeyboardWExtra is a function that creates a reply keyboard from set of parameters with extra capabilities
func (handler *TelegramBotHandler) CreateReplyKeyboardWExtra(resizeKeyboard, oneTimeKeyboard bool, keyboardButtons ...[]entity.ReplyKeyboardButton) string {

	buttonRows := make([][]*entity.ReplyKeyboardButton, 0)

	for _, keyboardRow := range keyboardButtons {

		row := make([]*entity.ReplyKeyboardButton, 0)
		for _, keyboardButton := range keyboardRow {
			button := new(entity.ReplyKeyboardButton)
			button.Text = keyboardButton.Text
			button.RequestContact = keyboardButton.RequestContact
			row = append(row, button)
		}

		buttonRows = append(buttonRows, row)
	}

	keyboard := entity.ReplyKeyboardMarkup{
		Keyboard:        buttonRows,
		ResizeKeyboard:  resizeKeyboard,
		OneTimeKeyboard: oneTimeKeyboard,
	}

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}

// CreateInlineKeyboard is a function that creates an inline keyboard from set of parameters for a chat
func (handler *TelegramBotHandler) CreateInlineKeyboard(keyboardButtons ...[]entity.InlineKeyboardButton) string {

	buttonRows := make([][]*entity.InlineKeyboardButton, 0)

	for _, keyboardRow := range keyboardButtons {

		row := make([]*entity.InlineKeyboardButton, 0)
		for _, keyboardButton := range keyboardRow {
			button := new(entity.InlineKeyboardButton)
			button.Text = keyboardButton.Text
			button.URL = keyboardButton.URL
			button.CallbackData = keyboardButton.CallbackData
			row = append(row, button)
		}

		buttonRows = append(buttonRows, row)
	}

	keyboard := entity.InlineKeyboardMarkup{
		InlineKeyboard: buttonRows,
	}

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}
