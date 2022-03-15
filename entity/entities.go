package entity

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateID          int64         `json:"update_id"`
	Message           Message       `json:"message"`
	EditedMessage     Message       `json:"edited_message"`
	ChannelPost       Message       `json:"channel_post"`
	EditedChannelPost Message       `json:"edited_channel_post"`
	CallbackQuery     CallbackQuery `json:"callback_query"`
	// 	InlineQuery        InlineQuery        `json:"inline_query"`
	// 	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	// 	ShippingQuery      ShippingQuery      `json:"shipping_query"`
	// 	PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
	// 	Poll               Poll               `json:"poll"`
	// 	PollAnswer         PollAnswer         `json:"poll_answer"`
	// 	ChatMemberUpdated  ChatMemberUpdated  `json:"my_chat_member"`
	// 	ChatMemberUpdated  ChatMemberUpdated  `json:"chat_member"`
}

// Message is a Telegram object that can be found inside an update.
type Message struct {
	MessageID            int64  `json:"message_id"`
	From                 User   `json:"from"`
	SenderChat           Chat   `json:"sender_chat"`
	Date                 int64  `json:"date"`
	Chat                 Chat   `json:"chat"`
	ForwardFrom          User   `json:"forward_from"`
	ForwardFromChat      Chat   `json:"forward_from_chat"`
	ForwardFromMessageID int64  `json:"forward_from_message_id"`
	ForwardSignature     string `json:"forward_signature"`
	ForwardSenderName    string `json:"forward_sender_name"`
	ForwardDate          int64  `json:"forward_date"`
	// ReplyToMessage        Message              `json:"reply_to_message"`
	ViaBot                User                 `json:"via_bot"`
	EditDate              int64                `json:"edit_date"`
	MediaGroupID          string               `json:"media_group_id"`
	AuthorSignature       string               `json:"author_signature"`
	Text                  string               `json:"text"`
	Document              Document             `json:"document"`
	Caption               string               `json:"caption"`
	Contact               Contact              `json:"contact"`
	NewChatMembers        []*User              `json:"new_chat_members"`
	LeftChatMember        User                 `json:"left_chat_member"`
	NewChatTitle          string               `json:"new_chat_title"`
	DeleteChatPhoto       bool                 `json:"delete_chat_photo"`
	GroupChatCreated      bool                 `json:"group_chat_created"`
	SuperGroupChatCreated bool                 `json:"supergroup_chat_created"`
	ChannelChatCreated    bool                 `json:"channel_chat_created"`
	MigrateToChatID       int64                `json:"migrate_to_chat_id"`
	MigrateFromChatID     int64                `json:"migrate_from_chat_id"`
	ConnectedWebsite      string               `json:"connected_website"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`
	Entities              []*MessageEntity     `json:"entities"`
	CaptionEntities       []*MessageEntity     `json:"caption_entities"`
	Animation             Animation            `json:"animation"`
	Video                 Video                `json:"video"`
	// Audio                         Audio                         `json:"audio"`
	// Photo                         []*PhotoSize                  `json:"photo"`
	// Sticker                       Sticker                       `json:"sticker"`
	// VideoNote                     VideoNote                     `json:"video_note"`
	// Voice                         Voice                         `json:"voice"`
	// PinnedMessage         		 Message              		   `json:"pinned_message"`
	// Dice                          Dice                          `json:"dice"`
	// Game                          Game                          `json:"game"`
	// Poll                          Poll                          `json:"poll"`
	// Venue                         Venue                         `json:"venue"`
	// Location                      Location                      `json:"location"`
	// new_chat_photo                []*PhotoSize                  `json:"new_chat_photo"`
	// MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	// Invoice                       Invoice                       `json:"invoice"`
	// SuccessfulPayment             SuccessfulPayment             `json:"successful_payment"`
	// PassportData                  PassportData                  `json:"passport_data"`
	// ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	// VoiceChatScheduled           VoiceChatScheduled             `json:"voice_chat_scheduled"`
	// VoiceChatStarted             VoiceChatStarted               `json:"voice_chat_started"`
	// VoiceChatEnded               VoiceChatEnded                  `json:"voice_chat_ended"`
	// VoiceChatParticipantsInvited VoiceChatParticipantsInvited    `json:"voice_chat_participants_invited"`

}

// CallbackQuery is a Telegram object that can be found inside an update.
type CallbackQuery struct {
	ID              string  `json:"id"`
	User            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"game_short_name"`
}

// MessageResponse is a response from a telegram bot after performing certain action like sending or editing message
type MessageResponse struct {
	Ok          bool    `json:"ok"`
	Result      Message `json:"result"`
	ErrorCode   int64   `json:"error_code"`
	Description string  `json:"description"`
}

// ChatResponse is a response from a telegram bot after performing certain action like getting chat
type ChatResponse struct {
	Ok          bool   `json:"ok"`
	Result      Chat   `json:"result"`
	ErrorCode   int64  `json:"error_code"`
	Description string `json:"description"`
}

// ChatMemberResponse is a response from a telegram bot after performing certain action like getting chat member
type ChatMemberResponse struct {
	Ok          bool       `json:"ok"`
	Result      ChatMember `json:"result"`
	ErrorCode   int64      `json:"error_code"`
	Description string     `json:"description"`
}

// ChatInviteLinkResponse is a response from a telegram bot after creating or editing chat invite link
type ChatInviteLinkResponse struct {
	Ok          bool           `json:"ok"`
	Result      ChatInviteLink `json:"result"`
	ErrorCode   int64          `json:"error_code"`
	Description string         `json:"description"`
}

// ChatDefaultResponse is a response from a telegram bot with no result value
type ChatDefaultResponse struct {
	Ok          bool        `json:"ok"`
	Result      interface{} `json:"result"`
	ErrorCode   int64       `json:"error_code"`
	Description string      `json:"description"`
}

// ChatMembersResponse is a response from a telegram bot after performing certain action like getting chat administrators
type ChatMembersResponse struct {
	Ok          bool         `json:"ok"`
	Result      []ChatMember `json:"result"`
	ErrorCode   int64        `json:"error_code"`
	Description string       `json:"description"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID                    int64  `json:"id"`
	Type                  string `json:"type"`
	Title                 string `json:"title"`
	UserName              string `json:"username"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Description           string `json:"description"`
	InviteLink            string `json:"invite_link"`
	PinnedMessage         string `json:"pinned_message"`
	SlowModeDelay         int64  `json:"slow_mode_delay"`
	MessageAutoDeleteTime int64  `json:"message_auto_delete_time"`
	StickerSetName        string `json:"sticker_set_name"`
	CanSetStickerSet      bool   `json:"can_set_sticker_set"`
	LinkedChatID          int64  `json:"linked_chat_id"`
	// Photo ChatPhoto   `json:"photo"`
	// Permissions           ChatPermissions `json:"permissions"`
	// Location ChatLocation   `json:"location"`
}

// User is a Telegram user object
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	UserName                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type ChatMember struct {
	// ChatMemberOwner
	Status      string `json:"status"`
	User        User   `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`

	// ChatMemberAdministrator
	CanBeEdited         bool `json:"can_be_edited"`
	CanMangeChat        bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVoiceChats bool `json:"can_manage_voice_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostMessages     bool `json:"can_post_messages"`
	CanEditMessages     bool `json:"can_edit_messages"`
	CanPinMessages      bool `json:"can_pin_messages"`

	// ChatMemberRestricted
	IsMember             bool  `json:"is_member"`
	CanSendMessages      bool  `json:"can_send_messages"`
	CanSendMediaMessages bool  `json:"can_send_media_messages"`
	CanSendPolls         bool  `json:"can_send_polls"`
	CanSendOtherMessages bool  `json:"can_send_other_messages"`
	CanAddWePagePreviews bool  `json:"can_add_web_page_previews"`
	UntilDate            int64 `json:"until_date"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 User   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name"`
	ExpireDate              int64  `json:"expire_date"`
	MemberLimit             int64  `json:"member_limit"`
	PendingJoinRequestCount int64  `json:"pending_join_request_count"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

// Document is a Telegram document object
type Document struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileName     string `json:"file_name"`
	MIMEType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
	// Thumb        PhotoSize `json:"thumb"`
}

// Document is a Telegram Video object
type Video struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	Duration     int64  `json:"duration"`
	FileName     string `json:"file_name"`
	MIMEType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
	// Thumb        PhotoSize `json:"thumb"`
}

// Animation is a Telegram animation object
type Animation struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	Duration     int64  `json:"duration"`
	FileName     string `json:"file_name"`
	MIMEType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
	// Thumb        PhotoSize `json:"thumb"`
}

// Contact is a Telegram contact object
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int64  `json:"user_id"`
	VCard       string `json:"vcard"`
}

// MessageEntity is a type that represents one special entity in a text message
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int64  `json:"offset"`
	Length   int64  `json:"length"`
	URL      string `json:"url"`
	User     User   `json:"user"`
	Language string `json:"language"`
}

// InputMediaPhoto is a type that represents a photo to be sent
type InputMediaPhoto struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
}

// InputMediaVideo is a type that represents a video to be sent
type InputMediaVideo struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	Width           int64            `json:"width"`
	Height          int64            `json:"height"`
	Duration        int64            `json:"duration"`
	SupportsString  bool             `json:"supports_streaming"`
}

// InputMediaAnimation is a type that represents a animation to be sent
type InputMediaAnimation struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	Width           int64            `json:"width"`
	Height          int64            `json:"height"`
	Duration        int64            `json:"duration"`
}

// InputMediaAudio is a type that represents a audio to be sent
type InputMediaAudio struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	Duration        int64            `json:"duration"`
	Performer       string           `json:"performer"`
	Title           string           `json:"title"`
}

// InputMediaDocument is a type that represents a general file to be sent
type InputMediaDocument struct {
	Type                        string           `json:"type"`
	Media                       string           `json:"media"`
	Thumb                       string           `json:"thumb"`
	Caption                     string           `json:"caption"`
	ParseMode                   string           `json:"parse_mode"`
	CaptionEntities             []*MessageEntity `json:"caption_entities"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"`
}

// Optional is a struct that holds all the optional values in one place
type Optional struct {
	ChatID                      interface{} // Since chat id can be integer or string
	MessageID                   int64
	InlineMessageID             string
	Text                        string
	ParseMode                   string
	Entities                    []*MessageEntity
	ReplyToMessageID            int64
	DisableNotification         bool
	DisableWebPageView          bool
	AllowSendingWithoutReply    bool
	ReplyMarkup                 string // It can be InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply so use the 'create' methods to create string representation of the methods
	ResizeKeyboard              bool
	OneTimeKeyboard             bool
	InputFieldPlaceholder       string
	Selective                   bool
	URL                         string
	ShowAlert                   bool
	CacheTime                   int64
	Duration                    int64
	Width                       int64
	Height                      int64
	ProtectContent              bool
	Thumb                       string // Only string thumb supported
	Caption                     string
	CaptionEntities             []*MessageEntity
	DisableContentTypeDetection bool

	// Invite link optional valus
	Name              string
	ExpireDate        int64
	MemberLimit       int64
	CreateJoinRequest bool

	// Chat member administration
	UntilDate           int64
	RevokeMessages      bool
	OnlyIfBanned        bool
	IsAnonymous         bool
	CanMangeChat        bool
	CanPostMessages     bool
	CanEditMessages     bool
	CanDeleteMessages   bool
	CanManageVoiceChats bool
	CanRestrictMembers  bool
	CanPromoteMembers   bool
	CanChangeInfo       bool
	CanInviteUsers      bool
	CanPinMessages      bool
}

// ReplyKeyboardMarkup is a struct that represents a reply to form Telegram keyboard
type ReplyKeyboardMarkup struct {
	Keyboard              [][]*ReplyKeyboardButton `json:"keyboard"`
	ResizeKeyboard        bool                     `json:"resize_keyboard"`
	OneTimeKeyboard       bool                     `json:"one_time_keyboard"`
	InputFieldPlaceholder string                   `json:"input_field_placeholder"`
	Selective             bool                     `json:"selective"`
}

// ReplyKeyboardButton is a struct that represents a Telegram reply keyboard button
type ReplyKeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
	// RequestPoll     KeyboardButtonPollType `json:"request_poll"`
}

// ReplyKeyboardRemove is a struct that represent a remove telegram reply keyboard command
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// ForceReply is a struct that enable clients to reply forcible
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
}

// InlineKeyboardMarkup is a struct that represents an inline keyboard for a reply chat
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is a struct that represents a Telegram inline keyboard button
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	URL                          string `json:"url"`
	CallbackData                 string `json:"callback_data"`
	SwitchInlineQuery            string `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`
	Pay                          bool   `json:"pay"`
	// LoginURL                     LoginURL     `json:"login_url"`
	// CallbackGame                 CallbackGame `json:"callback_game"`
}
