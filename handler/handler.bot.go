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
	optionals *entity.Optional) (*entity.MessageResponse, error) {

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

	botResponse := new(entity.MessageResponse)
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
	optionals *entity.Optional) (*entity.MessageResponse, error) {

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

	botResponse := new(entity.MessageResponse)
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
	optionals *entity.Optional) (*entity.MessageResponse, error) {

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
		replyMarkup = optionals.ReplyMarkup
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

	botResponse := new(entity.MessageResponse)
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

// SendVideoToTelegramChat sends a video files to the Telegram chat identified by its chat ID
/* Available Optional Values */
/* Duration                    int64 */
/* Width                       int64 */
/* Height                      int64 */
/* Thumb                       string */
/* Caption                     string */
/* ParseMode                   string */
/* CaptionEntities             []MessageEntity */
/* DisableContentTypeDetection bool */
/* DisableNotification         bool */
/* ReplyToMessageID            int64 */
/* ProtectContent              bool */
/* AllowSendingWithoutReply    bool */
/* ReplyMarkup                 string */
func (handler *TelegramBotHandler) SendVideoToTelegramChat(chatID interface{}, video string,
	optionals *entity.Optional) (*entity.MessageResponse, error) {

	caption := ""
	replyMarkup := ""
	parseMode := ""
	thumb := ""
	captionEntities := ""
	chatIDS := ""

	var duration int64
	var width int64
	var height int64
	var disableContentTypeDetection bool
	var disableNotification bool
	var replyToMessageID int64
	var allowSendingWithoutReply bool
	var protectContent bool

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

		duration = optionals.Duration
		width = optionals.Width
		height = optionals.Height
		thumb = optionals.Thumb
		caption = optionals.Caption
		parseMode = optionals.ParseMode
		disableContentTypeDetection = optionals.DisableContentTypeDetection
		disableNotification = optionals.DisableNotification
		replyToMessageID = optionals.ReplyToMessageID
		protectContent = optionals.ProtectContent
		allowSendingWithoutReply = optionals.AllowSendingWithoutReply
		replyMarkup = optionals.ReplyMarkup
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started sending video to telegram chat { Chat ID : %s, Video : %s, "+
		"Duration : %d, Width : %d, Height: %d, Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, "+
		"Disable Content Type Detection : %v, Disable Notification : %v, Reply To Message ID : %d, "+
		"Protect Content : %v, Allow Sending Without Reply : %v, Reply Markup : %s }",
		chatIDS, video, duration, width, height, thumb, caption, parseMode, captionEntities,
		disableContentTypeDetection, disableNotification, replyToMessageID, protectContent,
		allowSendingWithoutReply, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendVideo"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":                        {chatIDS},
			"video":                          {video},
			"duration":                       {strconv.FormatInt(duration, 10)},
			"width":                          {strconv.FormatInt(width, 10)},
			"height":                         {strconv.FormatInt(height, 10)},
			"thumb":                          {thumb},
			"caption":                        {caption},
			"parse_mode":                     {parseMode},
			"caption_entities":               {captionEntities},
			"disable_content_type_detection": {strconv.FormatBool(disableContentTypeDetection)},
			"disable_notification":           {strconv.FormatBool(disableNotification)},
			"protect_content":                {strconv.FormatBool(protectContent)},
			"reply_to_message_id":            {strconv.FormatInt(replyToMessageID, 10)},
			"allow_sending_without_reply":    {strconv.FormatBool(allowSendingWithoutReply)},
			"reply_markup":                   {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending video to telegram chat { Chat ID : %s, Video : %s, "+
			"Duration : %d, Width : %d, Height: %d, Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, "+
			"Disable Content Type Detection : %v, Disable Notification : %v, Reply To Message ID : %d, "+
			"Protect Content : %v, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, video, duration, width, height, thumb, caption, parseMode, captionEntities,
			disableContentTypeDetection, disableNotification, replyToMessageID, protectContent,
			allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.MessageResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending video to telegram chat, unable to parse response "+
			"{ Chat ID : %s, Video : %s, Duration : %d, Width : %d, Height: %d, Thumb : %s, Caption : %s, "+
			"Parse Mode : %s, Caption Entities : %s, Disable Content Type Detection : %v, Disable Notification : %v, "+
			"Reply To Message ID : %d, Protect Content : %v, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, video, duration, width, height, thumb, caption, parseMode, captionEntities,
			disableContentTypeDetection, disableNotification, replyToMessageID, protectContent,
			allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished sending video to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// SendAnimationToTelegramChat sends a animation files to the Telegram chat identified by its chat ID
/* Available Optional Values */
/* Duration                    int64 */
/* Width                       int64 */
/* Height                      int64 */
/* Thumb                       string */
/* Caption                     string */
/* ParseMode                   string */
/* CaptionEntities             []MessageEntity */
/* DisableContentTypeDetection bool */
/* DisableNotification         bool */
/* ReplyToMessageID            int64 */
/* ProtectContent              bool */
/* AllowSendingWithoutReply    bool */
/* ReplyMarkup                 string */
func (handler *TelegramBotHandler) SendAnimationToTelegramChat(chatID interface{}, animation string,
	optionals *entity.Optional) (*entity.MessageResponse, error) {

	caption := ""
	replyMarkup := ""
	parseMode := ""
	thumb := ""
	captionEntities := ""
	chatIDS := ""

	var duration int64
	var width int64
	var height int64
	var disableContentTypeDetection bool
	var disableNotification bool
	var replyToMessageID int64
	var allowSendingWithoutReply bool
	var protectContent bool

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

		duration = optionals.Duration
		width = optionals.Width
		height = optionals.Height
		thumb = optionals.Thumb
		caption = optionals.Caption
		parseMode = optionals.ParseMode
		disableContentTypeDetection = optionals.DisableContentTypeDetection
		disableNotification = optionals.DisableNotification
		replyToMessageID = optionals.ReplyToMessageID
		protectContent = optionals.ProtectContent
		allowSendingWithoutReply = optionals.AllowSendingWithoutReply
		replyMarkup = optionals.ReplyMarkup
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started sending animation to telegram chat { Chat ID : %s, Animation : %s, "+
		"Duration : %d, Width : %d, Height: %d, Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, "+
		"Disable Content Type Detection : %v, Disable Notification : %v, Reply To Message ID : %d, "+
		"Protect Content : %v, Allow Sending Without Reply : %v, Reply Markup : %s }",
		chatIDS, animation, duration, width, height, thumb, caption, parseMode, captionEntities,
		disableContentTypeDetection, disableNotification, replyToMessageID, protectContent,
		allowSendingWithoutReply, replyMarkup), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/sendAnimation"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":                        {chatIDS},
			"animation":                      {animation},
			"duration":                       {strconv.FormatInt(duration, 10)},
			"width":                          {strconv.FormatInt(width, 10)},
			"height":                         {strconv.FormatInt(height, 10)},
			"thumb":                          {thumb},
			"caption":                        {caption},
			"parse_mode":                     {parseMode},
			"caption_entities":               {captionEntities},
			"disable_content_type_detection": {strconv.FormatBool(disableContentTypeDetection)},
			"disable_notification":           {strconv.FormatBool(disableNotification)},
			"protect_content":                {strconv.FormatBool(protectContent)},
			"reply_to_message_id":            {strconv.FormatInt(replyToMessageID, 10)},
			"allow_sending_without_reply":    {strconv.FormatBool(allowSendingWithoutReply)},
			"reply_markup":                   {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending animation to telegram chat { Chat ID : %s, Animation : %s, "+
			"Duration : %d, Width : %d, Height: %d, Thumb : %s, Caption : %s, Parse Mode : %s, Caption Entities : %s, "+
			"Disable Content Type Detection : %v, Disable Notification : %v, Reply To Message ID : %d, "+
			"Protect Content : %v, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, animation, duration, width, height, thumb, caption, parseMode, captionEntities,
			disableContentTypeDetection, disableNotification, replyToMessageID, protectContent,
			allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.MessageResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For sending animation to telegram chat, unable to parse response "+
			"{ Chat ID : %s, Animation : %s, Duration : %d, Width : %d, Height: %d, Thumb : %s, Caption : %s, "+
			"Parse Mode : %s, Caption Entities : %s, Disable Content Type Detection : %v, Disable Notification : %v, "+
			"Reply To Message ID : %d, Protect Content : %v, Allow Sending Without Reply : %v, Reply Markup : %s }, %s",
			chatIDS, animation, duration, width, height, thumb, caption, parseMode, captionEntities,
			disableContentTypeDetection, disableNotification, replyToMessageID, protectContent,
			allowSendingWithoutReply, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished sending animation to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// EditMediaToTelegramChat edits a reply sent to the Telegram chat identified by its (chat ID and message ID) or inline message id
/* Only text is required because (chat ID and message ID) or inline message id are interchangable, if one is available it works */
/* Available Optional Values */
/* ChatID                   interface{} */
/* MessageID                int64 */
/* InlineMessageID          string */
/* ReplyMarkup              string */
func (handler *TelegramBotHandler) EditMediaToTelegramChat(media interface{},
	optionals *entity.Optional) (*entity.MessageResponse, error) {

	chatID := ""
	messageID := optionals.MessageID
	inlineMessageID := optionals.InlineMessageID
	replyMarkup := optionals.ReplyMarkup

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

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started editing media reply sent to telegram chat { Chat ID : %s, Message ID : %d, "+
		"Inline Message ID : %s, Media : %s, Reply Markup : %s }", chatID, messageID, inlineMessageID, media,
		replyMarkup), log.BotLogFile)

	mediaByte, _ := json.MarshalIndent(media, "", "	")
	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/editMessageMedia"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":           {chatID},
			"message_id":        {strconv.FormatInt(messageID, 10)},
			"inline_message_id": {inlineMessageID},
			"media":             {string(mediaByte)},
			"reply_markup":      {replyMarkup},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For editing media reply sent to telegram chat { Chat ID : %s, Message ID : %d, "+
			"Inline Message ID : %s, Media : %s, Reply Markup : %s }, %s", chatID, messageID, inlineMessageID, media,
			replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.MessageResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For editing media reply sent to telegram chat, unable to parse response "+
			"{ Chat ID : %s, Message ID : %d, Inline Message ID : %s, Media : %s, Reply Markup : %s }, %s",
			chatID, messageID, inlineMessageID, media, replyMarkup, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished editing media reply sent to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// GetChat gets  up to date information about the chat. Returns a Chat object on success.
func (handler *TelegramBotHandler) GetChat(chatID interface{}) (*entity.ChatResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started getting chat { Chat ID : %s }", chatIDS), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/getChat?chat_id=" + chatIDS
	response, err := http.Get(telegramAPI)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For getting chat { Chat ID : %s }, %s",
			chatIDS, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For getting chat, unable to parse response "+
			"{ Chat ID : %s }, %s", chatIDS, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished getting chat, Bot Response => %s", botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// GetChatMembers gets information about a member of a chat. Returns a ChatMember object on success.
func (handler *TelegramBotHandler) GetChatMembers(chatID interface{}, userID int64) (*entity.ChatMemberResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started getting chat members { Chat ID : %s, User ID : %d }", chatIDS, userID),
		log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/getChatMember?" +
		fmt.Sprintf("chat_id=%s&user_id=%d", chatIDS, userID)
	response, err := http.Get(telegramAPI)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For getting chat members { Chat ID : %s, User ID : %d }, %s",
			chatIDS, userID, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatMemberResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For getting chat members, unable to parse response "+
			"{ Chat ID : %s, User ID : %d }, %s", chatIDS, userID, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished getting chat members, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// GetChatAdministrators gets a list of administrators in a chat
func (handler *TelegramBotHandler) GetChatAdministrators(chatID interface{}) (*entity.ChatMembersResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started getting chat administrators { Chat ID : %s }", chatIDS),
		log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/getChatAdministrators?" +
		fmt.Sprintf("chat_id = %s", chatIDS)
	response, err := http.Get(telegramAPI)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For getting chat administrators { Chat ID : %s }, %s",
			chatIDS, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatMembersResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For getting chat administrators, unable to parse response "+
			"{ Chat ID : %s }, %s", chatIDS, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished getting chat administrators, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// ExportChatInviteLink generates a new primary invite link for a chat.
func (handler *TelegramBotHandler) ExportChatInviteLink(chatID interface{}) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started exporting chat invite link { Chat ID : %s }", chatIDS), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/exportChatInviteLink"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {chatIDS},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For exporting chat invite link { Chat ID : %s }, %s",
			chatIDS, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For exporting chat invite link, unable to parse response "+
			"{ Chat ID : %s }, %s", chatIDS, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished exporting chat invite link, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// CreateChatInviteLink creates an additional invite link for a chat
/* Available Optional Values */
/* Name                   string */
/* ExpireDate             int64 */
/* MemberLimit            int64 */
/* CreateJoinRequest      bool */
func (handler *TelegramBotHandler) CreateChatInviteLink(chatID interface{},
	optionals *entity.Optional) (*entity.ChatInviteLinkResponse, error) {

	chatIDS := ""

	var name string
	var expireDate int64
	var memberLimit int64
	var createsJoinRequest bool

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals aren't nil then set the values
	if optionals != nil {
		name = optionals.Name
		expireDate = optionals.ExpireDate
		memberLimit = optionals.MemberLimit
		createsJoinRequest = optionals.CreateJoinRequest
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started creating invite link to telegram chat { Chat ID : %s, Name : %s, "+
		"Expire Date : %d, Member Limit : %d, Create Join Request: %v }",
		chatIDS, name, expireDate, memberLimit, createsJoinRequest), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/createChatInviteLink"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":              {chatIDS},
			"name":                 {name},
			"expire_date":          {strconv.FormatInt(expireDate, 10)},
			"member_limit":         {strconv.FormatInt(memberLimit, 10)},
			"creates_join_request": {strconv.FormatBool(createsJoinRequest)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For creating invite link to telegram chat { Chat ID : %s, Name : %s, "+
			"Expire Date : %d, Member Limit : %d, Create Join Request: %v }, %s",
			chatIDS, name, expireDate, memberLimit, createsJoinRequest, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatInviteLinkResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For creating invite link to telegram chat, unable to parse response "+
			"{ Chat ID : %s, Name : %s, Expire Date : %d, Member Limit : %d, Create Join Request: %v }, %s",
			chatIDS, name, expireDate, memberLimit, createsJoinRequest, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished creating invite link to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// EditChatInviteLink edits a non-primary invite link created by the bot
/* Available Optional Values */
/* Name                   string */
/* ExpireDate             int64 */
/* MemberLimit            int64 */
/* CreateJoinRequest      bool */
func (handler *TelegramBotHandler) EditChatInviteLink(chatID interface{}, inviteLink string,
	optionals *entity.Optional) (*entity.ChatInviteLinkResponse, error) {

	chatIDS := ""

	var name string
	var expireDate int64
	var memberLimit int64
	var createsJoinRequest bool

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals aren't nil then set the values
	if optionals != nil {
		name = optionals.Name
		expireDate = optionals.ExpireDate
		memberLimit = optionals.MemberLimit
		createsJoinRequest = optionals.CreateJoinRequest
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started editing invite link to telegram chat { Chat ID : %s, Invite Link : %s, "+
		"Name : %s, Expire Date : %d, Member Limit : %d, Create Join Request: %v }",
		chatIDS, inviteLink, name, expireDate, memberLimit, createsJoinRequest), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/editChatInviteLink"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":              {chatIDS},
			"invite_link":          {inviteLink},
			"name":                 {name},
			"expire_date":          {strconv.FormatInt(expireDate, 10)},
			"member_limit":         {strconv.FormatInt(memberLimit, 10)},
			"creates_join_request": {strconv.FormatBool(createsJoinRequest)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For editing invite link to telegram chat { Chat ID : %s, Invite Link : %s, "+
			"Name : %s, Expire Date : %d, Member Limit : %d, Create Join Request: %v }, %s",
			chatIDS, inviteLink, name, expireDate, memberLimit, createsJoinRequest, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatInviteLinkResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For editing invite link to telegram chat, unable to parse response "+
			"{ Chat ID : %s, Invite Link : %s, Name : %s, Expire Date : %d, Member Limit : %d, "+
			"Create Join Request: %v }, %s", chatIDS, inviteLink, name, expireDate, memberLimit,
			createsJoinRequest, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished editing invite link to telegram chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// RevokeChatInviteLink revokes an invite link created by the bot.
func (handler *TelegramBotHandler) RevokeChatInviteLink(chatID interface{},
	inviteLink string) (*entity.ChatInviteLinkResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started revoking invite link { Chat ID : %s, Invite Link : %s }", chatIDS, inviteLink),
		log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/revokeChatInviteLink"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":     {chatIDS},
			"invite_link": {inviteLink},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For revoking invite link { Chat ID : %s, Invite Link : %s }, %s",
			chatIDS, inviteLink, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatInviteLinkResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For revoking invite link, unable to parse response "+
			"{ Chat ID : %s, Invite Link : %s }, %s", chatIDS, inviteLink, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished revoking invite link, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// ApproveChatJoinRequest approves a chat join request.
func (handler *TelegramBotHandler) ApproveChatJoinRequest(chatID interface{},
	userID int64) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started approving chat join request { Chat ID : %s, User ID : %d }", chatIDS, userID),
		log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/approveChatJoinRequest"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {chatIDS},
			"user_id": {strconv.FormatInt(userID, 10)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For approving chat join request { Chat ID : %s, User ID : %d }, %s",
			chatIDS, userID, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For approving chat join request, unable to parse response "+
			"{ Chat ID : %s, User ID : %d }, %s", chatIDS, userID, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished approving chat join request, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// DeclineChatJoinRequest declines a chat join request.
func (handler *TelegramBotHandler) DeclineChatJoinRequest(chatID interface{},
	userID int64) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started declining chat join request { Chat ID : %s, User ID : %d }", chatIDS, userID),
		log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/declineChatJoinRequest"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {chatIDS},
			"user_id": {strconv.FormatInt(userID, 10)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For declining chat join request { Chat ID : %s, User ID : %d }, %s",
			chatIDS, userID, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For declining chat join request, unable to parse response "+
			"{ Chat ID : %s, User ID : %d }, %s", chatIDS, userID, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished declining chat join request, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// BanChatMember bans a user in a group, a supergroup or a channel.
/* Available Optional Values */
/* UntilDate                  int64 */
/* RevokeMessages             bool */
func (handler *TelegramBotHandler) BanChatMember(chatID interface{}, userID int64,
	optionals *entity.Optional) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	var untilDate int64
	var revokeMessages bool

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals aren't nil then set the values
	if optionals != nil {
		untilDate = optionals.UntilDate
		revokeMessages = optionals.RevokeMessages
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started banning chat member { Chat ID : %s, User ID : %d, "+
		"Until Date : %d, Revoke Messages : %v }", chatIDS, userID, untilDate, revokeMessages), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/banChatMember"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":         {chatIDS},
			"user_id":         {strconv.FormatInt(userID, 10)},
			"until_date":      {strconv.FormatInt(untilDate, 10)},
			"revoke_messages": {strconv.FormatBool(revokeMessages)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For banning chat member { Chat ID : %s, User ID : %d, Until Date : %d, "+
			"Revoke Messages : %v }, %s", chatIDS, userID, untilDate, revokeMessages, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For banning chat member, unable to parse response "+
			"{ Chat ID : %s, User ID : %d, Until Date : %d, Revoke Messages : %v }, %s",
			chatIDS, userID, untilDate, revokeMessages, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished banning chat member, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// UnBanChatMember unban a previously banned user in a supergroup or channel.
/* Available Optional Values */
/* OnlyIfBanned             bool */
func (handler *TelegramBotHandler) UnbanChatMember(chatID interface{}, userID int64,
	optionals *entity.Optional) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	var onlyIfBanned bool

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals aren't nil then set the values
	if optionals != nil {
		onlyIfBanned = optionals.OnlyIfBanned
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started unbanning chat member { Chat ID : %s, User ID : %d, Only If Banned : %v }",
		chatIDS, userID, onlyIfBanned), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/unbanChatMember"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":        {chatIDS},
			"user_id":        {strconv.FormatInt(userID, 10)},
			"only_if_banned": {strconv.FormatBool(onlyIfBanned)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For unbanning chat member { Chat ID : %s, User ID : %d, Only If Banned : %v }, %s",
			chatIDS, userID, onlyIfBanned, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For unbanning chat member, unable to parse response "+
			"{ Chat ID : %s, User ID : %d, Only If Banned : %v }, %s",
			chatIDS, userID, onlyIfBanned, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished unbanning chat member, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// RestrictChatMember restricts a user in a supergroup.
/* Available Optional Values */
/* UntilDate                  int64 */
func (handler *TelegramBotHandler) RestrictChatMember(chatID interface{}, userID int64, permissions *entity.ChatPermissions,
	optionals *entity.Optional) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	var untilDate int64

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	if permissions == nil {
		return nil, errors.New("permissions are required")
	}

	// If optionals aren't nil then set the values
	if optionals != nil {
		untilDate = optionals.UntilDate
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started restricting chat member { Chat ID : %s, User ID : %d, "+
		"Until Date : %d, Permissions : %s }", chatIDS, userID, untilDate, permissions.ToString()), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/restrictChatMember"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":     {chatIDS},
			"user_id":     {strconv.FormatInt(userID, 10)},
			"permissions": {permissions.ToString()},
			"until_date":  {strconv.FormatInt(untilDate, 10)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For restricting chat member { Chat ID : %s, User ID : %d, Until Date : %d, "+
			"Permissions : %s }, %s", chatIDS, userID, untilDate, permissions.ToString(), err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For restricting chat member, unable to parse response "+
			"{ Chat ID : %s, User ID : %d, Until Date : %d, Permissions : %s }, %s",
			chatIDS, userID, untilDate, permissions.ToString(), err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished restricting chat member, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// PromoteChatMember promote or demote a user in a supergroup or a channel.
/* Available Optional Values */
/* IsAnonymous                     bool */
/* CanMangeChat                    bool */
/* CanPostMessages                 bool */
/* CanEditMessages                 bool */
/* CanDeleteMessages               bool */
/* CanManageVoiceChats             bool */
/* CanRestrictMembers              bool */
/* CanPromoteMembers               bool */
/* CanChangeInfo                   bool */
/* CanInviteUsers                  bool */
/* CanPinMessages                  bool */
func (handler *TelegramBotHandler) PromoteChatMember(chatID interface{}, userID int64,
	optionals *entity.Optional) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	var isAnonymous bool
	var canMangeChat bool
	var canPostMessages bool
	var canEditMessages bool
	var canDeleteMessages bool
	var canManageVoiceChats bool
	var canRestrictMembers bool
	var canPromoteMembers bool
	var canChangeInfo bool
	var canInviteUsers bool
	var canPinMessages bool

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	// If optionals aren't nil then set the values
	if optionals != nil {
		isAnonymous = optionals.IsAnonymous
		canMangeChat = optionals.CanMangeChat
		canPostMessages = optionals.CanPostMessages
		canEditMessages = optionals.CanEditMessages
		canDeleteMessages = optionals.CanDeleteMessages
		canManageVoiceChats = optionals.CanManageVoiceChats
		canRestrictMembers = optionals.CanRestrictMembers
		canPromoteMembers = optionals.CanPromoteMembers
		canChangeInfo = optionals.CanChangeInfo
		canInviteUsers = optionals.CanInviteUsers
		canPinMessages = optionals.CanPinMessages
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started promoting chat member { Chat ID : %s, User ID : %d, "+
		"IsAnonymous : %v, CanMangeChat : %v, CanPostMessages : %v, CanEditMessages : %v, CanDeleteMessages : %v, "+
		"CanManageVoiceChats : %v, CanRestrictMembers : %v, CanPromoteMembers : %v, CanChangeInfo : %v, "+
		"CanInviteUsers : %v, CanPinMessages : %v }", chatIDS, userID, isAnonymous, canMangeChat, canPostMessages,
		canEditMessages, canDeleteMessages, canManageVoiceChats, canRestrictMembers, canPromoteMembers, canChangeInfo,
		canInviteUsers, canPinMessages), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/promoteChatMember"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":                {chatIDS},
			"user_id":                {strconv.FormatInt(userID, 10)},
			"is_anonymous":           {strconv.FormatBool(isAnonymous)},
			"can_manage_chat":        {strconv.FormatBool(canMangeChat)},
			"can_post_messages":      {strconv.FormatBool(canPostMessages)},
			"can_edit_messages":      {strconv.FormatBool(canEditMessages)},
			"can_delete_messages":    {strconv.FormatBool(canDeleteMessages)},
			"can_manage_voice_chats": {strconv.FormatBool(canManageVoiceChats)},
			"can_restrict_members":   {strconv.FormatBool(canRestrictMembers)},
			"can_promote_members":    {strconv.FormatBool(canPromoteMembers)},
			"can_change_info":        {strconv.FormatBool(canChangeInfo)},
			"can_invite_users":       {strconv.FormatBool(canInviteUsers)},
			"can_pin_messages":       {strconv.FormatBool(canPinMessages)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For promoting chat member { Chat ID : %s, User ID : %d, "+
			"IsAnonymous : %v, CanMangeChat : %v, CanPostMessages : %v, CanEditMessages : %v, CanDeleteMessages : %v, "+
			"CanManageVoiceChats : %v, CanRestrictMembers : %v, CanPromoteMembers : %v, CanChangeInfo : %v, "+
			"CanInviteUsers : %v, CanPinMessages : %v }, %s", chatIDS, userID, isAnonymous, canMangeChat, canPostMessages,
			canEditMessages, canDeleteMessages, canManageVoiceChats, canRestrictMembers, canPromoteMembers, canChangeInfo,
			canInviteUsers, canPinMessages, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For promoting chat member, unable to parse response { Chat ID : %s, User ID : %d, "+
			"IsAnonymous : %v, CanMangeChat : %v, CanPostMessages : %v, CanEditMessages : %v, CanDeleteMessages : %v, "+
			"CanManageVoiceChats : %v, CanRestrictMembers : %v, CanPromoteMembers : %v, CanChangeInfo : %v, "+
			"CanInviteUsers : %v, CanPinMessages : %v }, %s", chatIDS, userID, isAnonymous, canMangeChat, canPostMessages,
			canEditMessages, canDeleteMessages, canManageVoiceChats, canRestrictMembers, canPromoteMembers, canChangeInfo,
			canInviteUsers, canPinMessages, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished promoting chat member, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// SetChatAdministratorCustomTitle sets a custom title for an administrator in a supergroup promoted by the bot.
func (handler *TelegramBotHandler) SetChatAdministratorCustomTitle(chatID interface{}, userID int64,
	customTitle string) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""
	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started setting chat administrator's custom title { Chat ID : %s, User ID : %d, "+
		"Custom Title : %s }", chatIDS, userID, customTitle), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/setChatAdministratorCustomTitle"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":      {chatIDS},
			"user_id":      {strconv.FormatInt(userID, 10)},
			"custom_title": {customTitle},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For setting chat administrator's custom title { Chat ID : %s, User ID : %d, "+
			"Custom Title : %s }, %s", chatIDS, userID, customTitle, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For setting chat administrator's custom title, unable to parse response "+
			"{ Chat ID : %s, User ID : %d, Custom Title : %s }, %s", chatIDS, userID, customTitle, err.Error()),
			log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished setting chat administrator's custom title, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// BanChatSenderChat bans a channel chat in a supergroup or a channel.
func (handler *TelegramBotHandler) BanChatSenderChat(chatID interface{},
	senderChatId int64) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started banning chat sender chat { Chat ID : %s, Sender Chat ID : %d }",
		chatIDS, senderChatId), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/banChatSenderChat"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":        {chatIDS},
			"sender_chat_id": {strconv.FormatInt(senderChatId, 10)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For banning chat sender chat { Chat ID : %s, Sender Chat ID : %d }, %s",
			chatIDS, senderChatId, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For banning chat sender chat, unable to parse response "+
			"{ Chat ID : %s, Sender Chat ID : %d }, %s", chatIDS, senderChatId, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished banning chat sender chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// UnbanChatSenderChat unban a previously banned channel chat in a supergroup or channel.
func (handler *TelegramBotHandler) UnbanChatSenderChat(chatID interface{},
	senderChatId int64) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started unbanning chat sender chat { Chat ID : %s, Sender Chat ID : %d }",
		chatIDS, senderChatId), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/unbanChatSenderChat"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":        {chatIDS},
			"sender_chat_id": {strconv.FormatInt(senderChatId, 10)},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For unbanning chat sender chat { Chat ID : %s, Sender Chat ID : %d }, %s",
			chatIDS, senderChatId, err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For unbanning chat sender chat, unable to parse response "+
			"{ Chat ID : %s, Sender Chat ID : %d }, %s", chatIDS, senderChatId, err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished unbanning chat sender chat, Bot Response => %s",
		botResponse.ToString()), log.BotLogFile)

	return botResponse, nil
}

// SetChatPermissions sets default chat permissions for all members.
func (handler *TelegramBotHandler) SetChatPermissions(chatID interface{},
	permissions *entity.ChatPermissions) (*entity.ChatDefaultResponse, error) {

	chatIDS := ""

	if id, ok := chatID.(int64); ok {
		chatIDS = strconv.FormatInt(id, 10)
	} else if id, ok := chatID.(string); ok {
		chatIDS = id
	} else {
		return nil, errors.New("chat id can only be type string or integer")
	}

	if permissions == nil {
		return nil, errors.New("permissions are required")
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Started setting chat permissions { Chat ID : %s, Permissions : %s }",
		chatIDS, permissions.ToString()), log.BotLogFile)

	var telegramAPI string = handler.BotAPIAccessPoint + handler.BotAccessToken + "/setChatPermissions"
	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id":     {chatIDS},
			"permissions": {permissions.ToString()},
		})

	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For setting chat permissions { Chat ID : %s, Permissions : %s }, %s",
			chatIDS, permissions.ToString(), err.Error()), log.ErrorLogFile)

		return nil, err
	}
	defer response.Body.Close()

	botResponse := new(entity.ChatDefaultResponse)
	err = json.NewDecoder(response.Body).Decode(botResponse)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.Logging(fmt.Sprintf("Error: For setting chat permissions, unable to parse response "+
			"{ Chat ID : %s, Permissions : %s }, %s", chatIDS, permissions.ToString(), err.Error()), log.ErrorLogFile)

		return nil, err
	}

	/* ---------------------------- Logging ---------------------------- */
	handler.Logging(fmt.Sprintf("Finished setting chat permissions, Bot Response => %s",
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
	optionals *entity.Optional) (*entity.MessageResponse, error) {

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

	botResponse := new(entity.MessageResponse)
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
