package entity

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateID      int64         `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

// Message is a Telegram object that can be found inside an update.
type Message struct {
	Text     string    `json:"text"`
	Chat     Chat      `json:"chat"`
	User     TUser     `json:"from"`
	Document TDocument `json:"document"`
	Contact  TContact  `json:"contact"`
}

// CallbackQuery is a Telegram object that can be found inside an update.
type CallbackQuery struct {
	ID   string `json:"id"`
	Data string `json:"data"`
	User TUser  `json:"from"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID int64 `json:"id"`
}

// TUser is a Telegram user object
type TUser struct {
	ID           int64  `json:"id"`
	LanguageCode string `json:"language_code"`
}

// TDocument is a Telegram document object
type TDocument struct {
	ID       string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Name     string `json:"file_name"`
	Type     string `json:"mime_type"`
}

// TContact is a Telegram contact object
type TContact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// ReplyKeyboardMarkup is a struct that represents a reply to form Telegram keyboard
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*ReplyKeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                     `json:"resize_keyboard"`
	OneTimeKeyboard bool                     `json:"one_time_keyboard"`
}

// ReplyKeyboardButton is a struct that represents a Telegram reply keyboard button
type ReplyKeyboardButton struct {
	Text           string `json:"text"`
	RequestContact bool   `json:"request_contact"`
}

// ReplyKeyboardRemove is a struct that represent a remove telegram reply keyboard command
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// InlineKeyboardMarkup is a struct that represents an inline keyboard for a reply chat
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is a struct that represents a Telegram inline keyboard button
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url"`
	CallbackData string `json:"callback_data"`
}
