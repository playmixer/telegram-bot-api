package telegram_bot

type User struct {
}

type ChatType string

const (
	PRIVATE    ChatType = "private"
	GROUP      ChatType = "group"
	SUPERGROUP ChatType = "supergroup"
	CHANNEL    ChatType = "channel"
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
	Result_id         string   `json:"result_id"`         //The unique identifier for the result that was chosen
	From              User     `json:"from"`              //The user that chose the result
	Location          Location `json:"location"`          //Опционально. Sender location, only for bots that require user location
	Inline_message_id string   `json:"inline_message_id"` //Опционально. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query             string   `json:"query"`             //The query that was used to obtain the result
}

type CallbackQuery struct {
	Id                string  `json:"id"`                //Уникальный идентификатор запроса
	From              User    `json:"from"`              //Отправитель
	Message           Message `json:"message"`           //Опционально. Сообщение, к которому была привязана вызвавшая запрос кнопка. Обратите внимание: если сообщение слишком старое, содержание сообщения и дата отправки будут недоступны.
	Inline_message_id string  `json:"inline_message_id"` //Опционально. Идентификатор сообщения, отправленного через вашего бота во встроенном режиме
	Data              string  `json:"data"`              //Данные, связанные с кнопкой. Обратите внимание, что клиенты могут добавлять свои данные в это поле.
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
