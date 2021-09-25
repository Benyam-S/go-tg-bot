package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Benyam-S/go-tg-bot/entity"
	"github.com/Benyam-S/go-tg-bot/log"
)

// SendReplyToTelegramChat sends a reply to the Telegram chat identified by its chat ID
/* For removing the reply keyboard, use ReplyKeyboardRemove{RemoveKeyboard: true} object as 'Reply Markup' */
/* It can be used to sending chat to telegram channel by using channel name as chatID */
/* Available Optional Values */
/* ParseMode                string -- 'html' if not provided */
/* Entities                 []*MessageEntity */
/* ReplyToMessageID         int64 */
/* DisableNotification      bool */
/* DisableWebPageView       bool */
/* AllowSendingWithoutReply bool */
/* ReplyMarkup              string */
func (handler *TelegramBotHandler) SendReplyToTelegramChat(chatID interface{}, text string,
	optionals *entity.Optional) (*entity.BotResponse, error) {

	chatIDS := ""
	parseMode := ""
	entities := ""
	replyMarkup := ""

	var replyToMessageID int64
	var disableNotification bool
	var disableWebPageView bool
	var allowSendingWithoutReply bool

	// Checking the chatID type
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals are nil then set the default mode
	if optionals == nil {
		parseMode = "html"
	} else {
		if len(optionals.Entities) > 0 {
			entitiesByte, _ := json.Marshal(optionals.Entities)
			entities = string(entitiesByte)
		}

		// If parse mode is not provided set to 'html' by default
		if optionals.ParseMode == "" {
			parseMode = "html"
		} else {
			parseMode = optionals.ParseMode
		}

		replyToMessageID = optionals.ReplyToMessageID
		replyMarkup = optionals.ReplyMarkup
		disableNotification = optionals.DisableNotification
		disableWebPageView = optionals.DisableWebPageView
		allowSendingWithoutReply = optionals.AllowSendingWithoutReply
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started sending reply to telegram chat { Chat ID : %s, Text : %s, Parse Mode : %s, "+
		"Entities : %s, Reply To Message ID : %d, Disable Notification : %v, Disable Web Page Preview : %v, "+
		"Allow Sending Without Reply : %v, Reply Markup : %s }", chatIDS, text, parseMode, entities,
		replyToMessageID, disableNotification, disableWebPageView,
		allowSendingWithoutReply, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendMessage"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":                     {chatIDS},
			"text":                        {text},
			"parse_mode":                  {parseMode},
			"entities":                    {entities},
			"reply_to_message_id":         {strconv.FormatInt(replyToMessageID, 10)},
			"disable_notification":        {strconv.FormatBool(disableNotification)},
			"disable_web_page_preview":    {strconv.FormatBool(disableWebPageView)},
			"allow_sending_without_reply": {strconv.FormatBool(allowSendingWithoutReply)},
			"reply_markup":                {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending reply to telegram chat { Chat ID : %s, Text : %s, Parse Mode : %s, "+
			"Entities : %s, Reply To Message ID : %d, Disable Notification : %v, Disable Web Page Preview : %v, "+
			"Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, text, parseMode, entities, replyToMessageID, disableNotification, disableWebPageView,
			allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.BotResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending reply to telegram chat, unable to parse response "+
			"{ Chat ID : %s, Text : %s, Parse Mode : %s, Entities : %s, Reply To Message ID : %d, "+
			"Disable Notification : %v, Disable Web Page Preview : %v, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, text, parseMode, entities, replyToMessageID, disableNotification, disableWebPageView,
			allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished sending reply to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// EditReplyToTelegramChat edits a reply sent to the Telegram chat identified by its (chat ID and message ID) or inline message id
/* Only text is required because (chat ID and message ID) or inline message id are interchangable, if one is available it works */
/* Available Optional Values */
/* ChatID                   interface{} */
/* MessageID                int64 */
/* InlineMessageID          string */
/* ParseMode                string */
/* Entities                 []*MessageEntity */
/* DisableWebPageView       bool */
/* ReplyMarkup              string */
func (handler *TelegramBotHandler) EditReplyToTelegramChat(text string,
	optionals *entity.Optional) (*entity.BotResponse, error) {

	chatID := ""
	parseMode := ""
	entities := ""
	replyMarkup := ""
	inlineMessageID := ""

	var messageID int64
	var disableWebPageView bool

	// If optionals are nil then set the default mode
	if optionals == nil {
		parseMode = "html"
	} else {
		if id, ok := optionals.ChatID.(int64); ok {
			chatID = strconv.FormatInt(id, 10)
		} else if id, ok := optionals.ChatID.(string); ok {
			chatID = id
		} else if optionals.ChatID == nil {
			// Since chatID can be empty
			chatID = ""
		} else {
			return nil, errors.New("chat id can only be type string or integer")
		}

		if len(optionals.Entities) > 0 {
			entitiesByte, _ := json.Marshal(optionals.Entities)
			entities = string(entitiesByte)
		}

		// If parse mode is not provided set to 'html' by default
		if optionals.ParseMode == "" {
			parseMode = "html"
		} else {
			parseMode = optionals.ParseMode
		}

		messageID = optionals.MessageID
		inlineMessageID = optionals.InlineMessageID
		replyMarkup = optionals.ReplyMarkup
		disableWebPageView = optionals.DisableWebPageView
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started editing reply sent to telegram chat { Chat ID : %s, Message ID : %d, "+
		"Inline Message ID : %s, Text : %s, Parse Mode : %s, Entities : %s, Disable Web Page Preview : %v, "+
		"Reply Markup : %s }", chatID, messageID, inlineMessageID, text, parseMode,
		entities, disableWebPageView, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/editMessageText"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":                  {chatID},
			"message_id":               {strconv.FormatInt(messageID, 10)},
			"inline_message_id":        {inlineMessageID},
			"text":                     {text},
			"parse_mode":               {parseMode},
			"entities":                 {entities},
			"disable_web_page_preview": {strconv.FormatBool(disableWebPageView)},
			"reply_markup":             {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For editing reply sent to telegram chat { Chat ID : %s, Message ID : %d, "+
			"Inline Message ID : %s, Text : %s, Parse Mode : %s, Entities : %s, Disable Web Page Preview : %v, "+
			"Reply Markup : %s }, %s", chatID, messageID, inlineMessageID, text, parseMode,
			entities, disableWebPageView, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.BotResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For editing reply sent to telegram chat, unable to parse response { Chat ID : %s, Message ID : %d, "+
			"Inline Message ID : %s, Text : %s, Parse Mode : %s, Entities : %s, Disable Web Page Preview : %v, "+
			"Reply Markup : %s }, %s", chatID, messageID, inlineMessageID, text, parseMode,
			entities, disableWebPageView, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished editing reply sent to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// SendDocumentToTelegramChat sends a document to the Telegram chat identified by its chat ID
/* Available Optional Values */
/* Thumb                       string */
/* Caption                     string */
/* ParseMode                   string */
/* CaptionEntities             []MessageEntity */
/* DisableContentTypeDetection bool */
/* DisableNotification         bool */
/* ReplyToMessageID            int64 */
/* AllowSendingWithoutReply    bool */
/* ReplyMarkup                 string */
func (handler *TelegramBotHandler) SendDocumentToTelegramChat(chatID interface{}, fileID string,
	optionals *entity.Optional) (*entity.BotResponse, error) {

	caption := ""
	replyMarkup := ""
	parseMode := ""
	thumb := ""
	captionEntities := ""
	chatIDS := ""

	var disableContentTypeDetection bool
	var disableNotification bool
	var replyToMessageID int64
	var allowSendingWithoutReply bool

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals are nil then set the default mode
	if optionals == nil {
		parseMode = "html"
	} else {
		if len(optionals.CaptionEntities) > 0 {
			captionEntitiesByte, _ := json.Marshal(optionals.CaptionEntities)
			captionEntities = string(captionEntitiesByte)
		}

		thumb = optionals.Thumb
		caption = optionals.Caption
		parseMode = optionals.ParseMode
		disableContentTypeDetection = optionals.DisableContentTypeDetection
		disableNotification = optionals.DisableNotification
		replyToMessageID = optionals.ReplyToMessageID
		allowSendingWithoutReply = optionals.AllowSendingWithoutReply
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started sending document to telegram chat { Chat ID : %s, Document : %s, "+
		"Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, Disable Content Type Detection : %v, "+
		"Disable Notification : %v, Reply To Message ID : %d, Allow Sending Without Reply : %v, Reply Markup : %s }",
		chatIDS, fileID, thumb, caption, parseMode, captionEntities, disableContentTypeDetection,
		disableNotification, replyToMessageID, allowSendingWithoutReply, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendDocument"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":                        {chatIDS},
			"document":                       {fileID},
			"thumb":                          {thumb},
			"caption":                        {caption},
			"parse_mode":                     {parseMode},
			"caption_entities":               {captionEntities},
			"disable_content_type_detection": {strconv.FormatBool(disableContentTypeDetection)},
			"disable_notification":           {strconv.FormatBool(disableNotification)},
			"reply_to_message_id":            {strconv.FormatInt(replyToMessageID, 10)},
			"allow_sending_without_reply":    {strconv.FormatBool(allowSendingWithoutReply)},
			"reply_markup":                   {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending document to telegram chat { Chat ID : %s, Document : %s, "+
			"Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, Disable Content Type Detection : %v, "+
			"Disable Notification : %v, Reply To Message ID : %d, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, fileID, thumb, caption, parseMode, captionEntities, disableContentTypeDetection,
			disableNotification, replyToMessageID, allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.BotResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending document to telegram chat, unable to parse response { Chat ID : %s, Document : %s, "+
			"Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, Disable Content Type Detection : %v, "+
			"Disable Notification : %v, Reply To Message ID : %d, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, fileID, thumb, caption, parseMode, captionEntities, disableContentTypeDetection,
			disableNotification, replyToMessageID, allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished sending document to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// AnswerToTelegramCallBack sends a reply to the Telegram call back request identified by the query id
/* Available Optional Values */
/* Text                     string */
/* URL                      string */
/* ShowAlert                bool */
/* CacheTime                int64 */
func (handler *TelegramBotHandler) AnswerToTelegramCallBack(queryID string,
	optionals *entity.Optional) (*entity.BotResponse, error) {

	text := ""
	callbackUrl := ""

	var showAlert bool
	var cacheTime int64

	if optionals != nil {
		text = optionals.Text
		callbackUrl = optionals.URL
		showAlert = optionals.ShowAlert
		cacheTime = optionals.CacheTime
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started answering to telegram callback { Callback Query ID : %s, Text : %s, "+
		"Show Alert : %v, URL : %s, Cache Time : %d }", queryID, text, showAlert, callbackUrl, cacheTime),
		log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/answerCallbackQuery"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"callback_query_id": {queryID},
			"text":              {text},
			"show_alert":        {strconv.FormatBool(showAlert)},
			"url":               {callbackUrl},
			"cache_time":        {strconv.FormatInt(cacheTime, 10)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For answering to telegram callback { Callback Query ID : %s, Text : %s, "+
			"Show Alert : %v, URL : %s, Cache Time : %d }, %s", queryID, text, showAlert, callbackUrl,
			cacheTime, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.BotResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For answering to telegram callback, unable to parse response  { Callback Query ID : %s, Text : %s, "+
			"Show Alert : %v, URL : %s, Cache Time : %d }, %s", queryID, text, showAlert, callbackUrl,
			cacheTime, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished answering to telegram callback, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// CreateReplyKeyboard is a function that creates a reply keyboard from set of parameters
/* ResizeKeyboard              bool -- True if not provided */
/* OneTimeKeyboard             bool */
/* Selective                   bool */
func (handler *TelegramBotHandler) CreateReplyKeyboard(optionals *entity.Optional,
	keyboardButtons ...[]*entity.ReplyKeyboardButton) string {

	var resizeKeyboard bool
	var oneTimeKeyboard bool
	var selective bool

	inputFieldPlaceholder := ""
	buttonRows := make([][]*entity.ReplyKeyboardButton, 0)

	for _, keyboardRow := range keyboardButtons {
		row := make([]*entity.ReplyKeyboardButton, 0)
		row = append(row, keyboardRow...)

		buttonRows = append(buttonRows, row)
	}

	// optionals are different from nil then set the values
	if optionals != nil {

		resizeKeyboard = optionals.ResizeKeyboard
		oneTimeKeyboard = optionals.OneTimeKeyboard
		inputFieldPlaceholder = optionals.InputFieldPlaceholder
		selective = optionals.Selective
	} else {
		// set default values
		resizeKeyboard = true
	}

	keyboard := entity.ReplyKeyboardMarkup{
		Keyboard:              buttonRows,
		ResizeKeyboard:        resizeKeyboard,
		OneTimeKeyboard:       oneTimeKeyboard,
		InputFieldPlaceholder: inputFieldPlaceholder,
		Selective:             selective,
	}

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}

// CreateInlineKeyboard is a function that creates an inline keyboard from set of parameters for a chat
func (handler *TelegramBotHandler) CreateInlineKeyboard(keyboardButtons ...[]*entity.InlineKeyboardButton) string {

	buttonRows := make([][]*entity.InlineKeyboardButton, 0)

	for _, keyboardRow := range keyboardButtons {
		row := make([]*entity.InlineKeyboardButton, 0)
		row = append(row, keyboardRow...)

		buttonRows = append(buttonRows, row)
	}

	keyboard := entity.InlineKeyboardMarkup{
		InlineKeyboard: buttonRows,
	}

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}

// CreateReplyKeyboardRemove is a function that creates a remove reply keyboard
func (handler *TelegramBotHandler) CreateReplyKeyboardRemove(keyboard *entity.ReplyKeyboardRemove) string {

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}

// CreateForceReplyKeyboard is a function that creates a force reply keyboard
func (handler *TelegramBotHandler) CreateForceReplyKeyboard(keyboard *entity.ForceReply) string {

	keyboardS, _ := json.Marshal(&keyboard)
	return string(keyboardS)
}
