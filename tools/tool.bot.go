package tools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/Benyam-S/go-tg-bot/entity"
)

// SendReplyToTelegramChat sends a reply to the Telegram chat identified by its chat Id
func SendReplyToTelegramChat(chatID int64, reply ...string) (string, error) {

	text := ""
	replyMarkup := ""

	if len(reply) > 0 {
		text = reply[0]
	}

	if len(reply) > 1 {
		replyMarkup = reply[1]
	}

	var telegramAPI string = os.Getenv("api_access_point") + os.Getenv("bot_api_token") + "/sendMessage"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":      {strconv.FormatInt(chatID, 10)},
			"text":         {text},
			"reply_markup": {replyMarkup},
			"parse_mode":   {"html"},
		})

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// SendDocumentToTelegramChat sends a document to the Telegram chat identified by its chat Id
func SendDocumentToTelegramChat(chatID int64, fileID string, reply ...string) (string, error) {

	caption := ""
	replyMarkup := ""

	if len(reply) > 0 {
		caption = reply[0]
	}

	if len(reply) > 1 {
		replyMarkup = reply[1]
	}

	var telegramAPI string = os.Getenv("api_access_point") + os.Getenv("bot_api_token") + "/sendDocument"
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
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// PostToTelegramChannel posts a certain content to a telegram channel
func PostToTelegramChannel(post ...string) (string, error) {

	text := ""
	replyMarkup := ""

	if len(post) > 0 {
		text = post[0]
	}

	if len(post) > 1 {
		replyMarkup = post[1]
	}

	var telegramAPI string = os.Getenv("api_access_point") + os.Getenv("bot_api_token") + "/sendMessage"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":      {os.Getenv("channel_name")},
			"text":         {text},
			"reply_markup": {replyMarkup},
			"parse_mode":   {"html"},
		})

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// AnswerToTelegramCallBack sends a reply to the Telegram call back request identified by the query id
func AnswerToTelegramCallBack(queryID string, text string) (string, error) {

	var telegramAPI string = os.Getenv("api_access_point") + os.Getenv("bot_api_token") + "/answerCallbackQuery"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"callback_query_id": {queryID},
			"text":              {text},
		})

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// CreateReplyKeyboard is a function that creates a reply keyboard from set of parameters
func CreateReplyKeyboard(resizeKeyboard, oneTimeKeyboard bool, keyboardButtons ...[]string) string {

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
func CreateReplyKeyboardWExtra(resizeKeyboard, oneTimeKeyboard bool, keyboardButtons ...[]entity.ReplyKeyboardButton) string {

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
func CreateInlineKeyboard(keyboardButtons ...[]entity.InlineKeyboardButton) string {

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
