package telegram_bot

type User struct {
	Id                      int32  `json:"id"`                          // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	IsBot                   bool   `json:"is_bot"`                      // True, if this user is a bot
	FirstName               string `json:"first_name"`                  // User's or bot's first name
	LastName                string `json:"last_name"`                   // Optional. User's or bot's last name
	Username                string `json:"username"`                    // Optional. User's or bot's username
	LanguageCode            string `json:"language_code"`               // Optional. IETF language tag of the user's language
	IsPremium               bool   `json:"is_premium"`                  // Optional. True, if this user is a Telegram Premium user
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`    // Optional. True, if this user added the bot to the attachment menu
	CanJoinGroups           bool   `json:"can_join_groups"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
	CanConnectToBusiness    bool   `json:"can_connect_to_business"`     // Optional. True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
}

type ChatType string

const (
	PRIVATE    ChatType = "private"
	GROUP      ChatType = "group"
	SUPERGROUP ChatType = "supergroup"
	CHANNEL    ChatType = "channel"
)

type ChatAction string

const (
	TYPING          ChatAction = "typing"
	UPLOAD_PHOTO    ChatAction = "upload_photo"
	RECORD_VIDEO    ChatAction = "record_video"
	UPLOAD_VIDEO    ChatAction = "upload_video"
	RECORD_AUDIO    ChatAction = "record_audio"
	UPLOAD_AUDIO    ChatAction = "upload_audio"
	UPLOAD_DOCUMENT ChatAction = "upload_document"
	FIND_LOCATION   ChatAction = "find_location"
)

type Chat struct {
	Id                          int64    `json:"id"`                             //Уникальный идентификатор чата. Абсолютное значение не превышает 1e13
	Type                        ChatType `json:"type"`                           //Тип чата: “private”, “group”, “supergroup” или “channel”
	Title                       string   `json:"title"`                          //Опционально. Название, для каналов или групп
	Username                    string   `json:"username"`                       //Опционально. Username, для чатов и некоторых каналов
	FirstName                   string   `json:"first_name"`                     //Опционально. Имя собеседника в чате
	LastName                    string   `json:"last_name"`                      //Опционально. Фамилия собеседника в чате
	AllMembersAreAdministrators bool     `json:"all_members_are_administrators"` //Опционально.True, если все участники чата являются администраторами
}

type MessageEntity struct {
}

type Audio struct {
}

type Document struct {
}

type PhotoSize struct {
}

type Sticker struct {
}

type Video struct {
}

type Voice struct {
}

type Contact struct {
}

type Location struct {
}

type Venue struct {
}

type ReplyToMessage struct {
	MessageId             int64           `json:"message_id"`
	From                  User            `json:"from"`
	Date                  int64           `json:"date"`
	Chat                  Chat            `json:"chat"`
	ForwardFrom           User            `json:"forward_from"`
	ForwardDate           int64           `json:"forward_date"`
	Text                  string          `json:"text"`
	Entities              []MessageEntity `json:"entities"`
	Audio                 Audio           `json:"audio"`                   //Опционально. Информация об аудиофайле
	Document              Document        `json:"document"`                // Опционально. Информация о файле
	Photo                 []PhotoSize     `json:"photo"`                   //Опционально. Доступные размеры фото
	Sticker               Sticker         `json:"sticker"`                 //Опционально. Информация о стикере
	Video                 Video           `json:"video"`                   //Опционально. Информация о видеозаписи
	Voice                 Voice           `json:"voice"`                   //Опционально. Информация о голосовом сообщении
	Caption               string          `json:"caption"`                 //Опционально. Подпись к файлу, фото или видео, 0-200 символов
	Contact               Contact         `json:"contact"`                 //Опционально. Информация об отправленном контакте
	Location              Location        `json:"location"`                //Опционально. Информация о местоположении
	Venue                 Venue           `json:"venue"`                   //Опционально. Информация о месте на карте
	NewChatMember         User            `json:"new_chat_member"`         //Опционально. Информация о пользователе, добавленном в группу
	LeftChatMember        User            `json:"left_chat_member"`        //Опционально. Информация о пользователе, удалённом из группы
	NewChatTitle          string          `json:"new_chat_title"`          //Опционально. Название группы было изменено на это поле
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo"`          //Опционально. Фото группы было изменено на это поле
	DeleteChatPhoto       bool            `json:"delete_chat_photo"`       //Опционально. Сервисное сообщение: фото группы было удалено
	GroupChatCreated      bool            `json:"group_chat_created"`      //Опционально. Сервисное сообщение: группа создана
	SupergroupChatCreated bool            `json:"supergroup_chat_created"` //Опционально. Сервисное сообщение: супергруппа создана
	ChannelChatCreated    bool            `json:"channel_chat_created"`    //Опционально. Сервисное сообщение: канал создан
	MigrateToChatId       int64           `json:"migrate_to_chat_id"`      //Опционально. Группа была преобразована в супергруппу с указанным идентификатором. Не превышает 1e13
	MigrateFromChatId     int64           `json:"migrate_from_chat_id"`    //Опционально. Cупергруппа была создана из группы с указанным идентификатором. Не превышает 1e13
}

type Message struct {
	MessageId             int64           `json:"message_id"`
	From                  User            `json:"from"`
	Date                  int64           `json:"date"`
	Chat                  Chat            `json:"chat"`
	ForwardFrom           User            `json:"forward_from"`
	ForwardDate           int64           `json:"forward_date"`
	ReplyToMessage        ReplyToMessage  `json:"reply_to_message"`
	Text                  string          `json:"text"`
	Entities              []MessageEntity `json:"entities"`
	Audio                 Audio           `json:"audio"`                   //Опционально. Информация об аудиофайле
	Document              Document        `json:"document"`                //Опционально. Информация о файле
	Photo                 []PhotoSize     `json:"photo"`                   //Опционально. Доступные размеры фото
	Sticker               Sticker         `json:"sticker"`                 //Опционально. Информация о стикере
	Video                 Video           `json:"video"`                   //Опционально. Информация о видеозаписи
	Voice                 Voice           `json:"voice"`                   //Опционально. Информация о голосовом сообщении
	Caption               string          `json:"caption"`                 //Опционально. Подпись к файлу, фото или видео, 0-200 символов
	Contact               Contact         `json:"contact"`                 //Опционально. Информация об отправленном контакте
	Location              Location        `json:"location"`                //Опционально. Информация о местоположении
	Venue                 Venue           `json:"venue"`                   //Опционально. Информация о месте на карте
	NewChatMember         User            `json:"new_chat_member"`         //Опционально. Информация о пользователе, добавленном в группу
	LeftChatMember        User            `json:"left_chat_member"`        //Опционально. Информация о пользователе, удалённом из группы
	NewChatTitle          string          `json:"new_chat_title"`          //Опционально. Название группы было изменено на это поле
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo"`          //Опционально. Фото группы было изменено на это поле
	DeleteChatPhoto       bool            `json:"delete_chat_photo"`       //Опционально. Сервисное сообщение: фото группы было удалено
	GroupChatCreated      bool            `json:"group_chat_created"`      //Опционально. Сервисное сообщение: группа создана
	SupergroupChatCreated bool            `json:"supergroup_chat_created"` //Опционально. Сервисное сообщение: супергруппа создана
	ChannelChatCreated    bool            `json:"channel_chat_created"`    //Опционально. Сервисное сообщение: канал создан
	MigrateToChatId       int64           `json:"migrate_to_chat_id"`      //Опционально. Группа была преобразована в супергруппу с указанным идентификатором. Не превышает 1e13
	MigrateFromChatId     int64           `json:"migrate_from_chat_id"`    //Опционально. Cупергруппа была создана из группы с указанным идентификатором. Не превышает 1e13
	PinnedMessage         ReplyToMessage  `json:"pinned_message"`          //Опционально. Указанное сообщение было прикреплено. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
}

type InlineQuery struct {
	Id       string   `json:"id"`       //Unique identifier for this query
	From     User     `json:"from"`     //Sender
	Location Location `json:"location"` //Опционально. Sender location, only for bots that request user location
	Query    string   `json:"query"`    //Text of the query
	Offset   string   `json:"offset"`   //Offset of the results to be returned, can be controlled by the bot
}

type ChosenInlineResult struct {
	Result_id       string   `json:"result_id"`         //The unique identifier for the result that was chosen
	From            User     `json:"from"`              //The user that chose the result
	Location        Location `json:"location"`          //Опционально. Sender location, only for bots that require user location
	InlineMessageId string   `json:"inline_message_id"` //Опционально. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string   `json:"query"`             //The query that was used to obtain the result
}

type CallbackQuery struct {
	Id              string  `json:"id"`                //Уникальный идентификатор запроса
	From            User    `json:"from"`              //Отправитель
	Message         Message `json:"message"`           //Опционально. Сообщение, к которому была привязана вызвавшая запрос кнопка. Обратите внимание: если сообщение слишком старое, содержание сообщения и дата отправки будут недоступны.
	InlineMessageId string  `json:"inline_message_id"` //Опционально. Идентификатор сообщения, отправленного через вашего бота во встроенном режиме
	Data            string  `json:"data"`              //Данные, связанные с кнопкой. Обратите внимание, что клиенты могут добавлять свои данные в это поле.
}

type UpdateResult struct {
	UpdateId           int64              `json:"update_id"`
	Message            Message            `json:"message"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      CallbackQuery      `json:"callback_query"`
}

type Update struct {
	Ok     bool           `json:"ok"`
	Result []UpdateResult `json:"result"`
}

type SendMessageResult struct {
	Ok          bool    `json:"ok"`
	Result      Message `json:"result"`
	ErrorCode   int     `json:"error_code"`
	Description string  `json:"description"`
}

type Result struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

// Формат сообщения
type MessageStyle string

const (
	MessageStyleMarkdownV2 MessageStyle = "MarkdownV2"
	MessageStyleHTML       MessageStyle = "HTML"
)

type MessageOptionsField string

const (
	MOFParseMode   MessageOptionsField = "parse_mode"
	MOFReplyMarkup MessageOptionsField = "reply_markup"
	// MOFInlineMarkup MessageOptionsField = "inline_keyboard"
)

type MessageOption struct {
	Field MessageOptionsField
	Value string
}

// markup keyboard

type KeyboardButtonRequestUsers struct {
	RequestId       int32 `json:"request_id"`                 // Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message
	UserIsBot       bool  `json:"user_is_bot,omitempty"`      // Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsPremium   bool  `json:"user_is_premium,omitempty"`  // Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	MaxQuantity     uint  `json:"max_quantity,omitempty"`     // Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     bool  `json:"request_name,omitempty"`     // Optional. Pass True to request the users' first and last name
	RequestUsername bool  `json:"request_username,omitempty"` // Optional. Pass True to request the users' username
	RequestPhoto    bool  `json:"request_photo,omitempty"`    // Optional. Pass True to request the users' photo
}

type ChatAdministratorRights struct {
	SAnonymous          bool `json:"s_anonymous"`                 // True, if the user's presence in the chat is hidden
	CanManageChat       bool `json:"can_manage_chat"`             // True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool `json:"can_delete_messages"`         // True, if the administrator can delete messages of other users
	CanManageVideoChats bool `json:"can_manage_video_chats"`      // True, if the administrator can manage video chats
	CanRestrictMembers  bool `json:"can_restrict_members"`        // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool `json:"can_promote_members"`         // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool `json:"can_change_info"`             // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool `json:"can_invite_users"`            // True, if the user is allowed to invite new users to the chat
	CanPostStories      bool `json:"can_post_stories"`            // True, if the administrator can post stories to the chat
	CanEditStories      bool `json:"can_edit_stories"`            // True, if the administrator can edit stories posted by other users
	CanDeleteStories    bool `json:"can_delete_stories"`          // True, if the administrator can delete stories posted by other users
	CanPostMessages     bool `json:"can_post_messages,omitempty"` // Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool `json:"can_edit_messages,omitempty"` // Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`  // Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanManageTopics     bool `json:"can_manage_topics,omitempty"` // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

type KeyboardButtonRequestChat struct {
	RequestId                int32                    `json:"request_id"`                          // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	ChatIsChannel            bool                     `json:"chat_is_channel"`                     // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsForum              bool                     `json:"chat_is_forum,omitempty"`             // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatHasUsername          bool                     `json:"chat_has_username,omitempty"`         // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatIsCreated            bool                     `json:"chat_is_created,omitempty"`           // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	UserAdministrator_rights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"` // Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministrator_rights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`  // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotIsMember              bool                     `json:"bot_is_member,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestTitle             bool                     `json:"request_title,omitempty"`             // Optional. Pass True to request the chat's title
	RequestUsername          bool                     `json:"request_username,omitempty"`          // Optional. Pass True to request the chat's username
	RequestPhoto             bool                     `json:"request_photo,omitempty"`             // Optional. Pass True to request the chat's photo
}

type KeyboardButtonPollType struct {
	Type string `json:"type"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

type WebAppInfo struct {
	Url string `json:"url"` // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}

type KeyboardButton struct {
	Text            string                      `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`    // Optional. If specified, pressing the button will open a list of suitable users. Identifiers of selected users will be sent to the bot in a “users_shared” service message. Available in private chats only.
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`     // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats only.
	RequestContact  bool                        `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestLocation bool                        `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`          // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a “web_app_data” service message. Available in private chats only.
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`                          // Array of button rows, each represented by an Array of KeyboardButton objects
	IsPersistent          bool               `json:"is_persistent,omitempty"`           // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             bool               `json:"selective,omitempty"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
}

// inline keyboard

type LoginUrl struct {
	Url                string `json:"url"`                  // An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.
	ForwardText        string `json:"forward_text"`         // Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query"`                         // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`    // Optional. True, if private chats with users can be chosen
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`     // Optional. True, if private chats with bots can be chosen
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`   // Optional. True, if group and supergroup chats can be chosen
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"` // Optional. True, if channel chats can be chosen
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button
	Url                          *string                      `json:"url,omitempty"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	CallbackData                 *string                      `json:"callback_data,omitempty"`                    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot.
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`                        // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted.
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field
	// CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional. Description of the game that will be launched when the user presses the button.
	Pay bool `json:"pay,omitempty"` // Optional. Specify True, to send a Pay button.
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`     // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
}

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`             // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder *string `json:"input_field_placeholder"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             bool    `json:"selective"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
}
