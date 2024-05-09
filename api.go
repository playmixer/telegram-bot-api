package telegram_bot

import (
	"fmt"
	"log"
	"net/url"
)

func (t *TelegramBot) GetUpdates(offset, limit, timeout int64) Update {
	if limit == 0 {
		limit = 100
	}
	url := t.GetApiUrl(fmt.Sprintf("%s?offset=%v&limit=%v&timeout=%v", "getUpdates", offset, limit, timeout))
	// fmt.Println(url)

	var data Update
	err := t.Get(url.String(), &data)
	if err != nil {
		return data
	}

	return data
}

type tURL string

func (u *tURL) Set(field MessageOptionsField, value string) error {
	_u, err := url.Parse(string(*u))
	if err != nil {
		return err
	}

	q := _u.Query()
	q.Set(string(field), value)

	_u.RawQuery = q.Encode()

	*u = tURL(_u.String())
	return nil
}

func (u *tURL) String() string {
	return string(*u)
}

type MessageOption func(u *tURL) error

func StyleMarkdown(v MessageStyle) MessageOption {
	return func(u *tURL) error {
		err := u.Set(MOFParseMode, string(v))
		return err
	}
}

func (t *TelegramBot) SendMessage(chatId int64, text string, options ...MessageOption) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("sendMessage?chat_id=%v&text=%s", chatId, formatingText(text)))

	for _, opt := range options {
		err := opt(&url)
		if err != nil {
			log.Printf("error option %s", err)
		}
	}

	log.Println(url)
	var result SendMessageResult
	err := t.Get(string(url), &result)
	if err != nil {
		return result
	}

	return result
}

func (t *TelegramBot) ReplyToMessage(chatId int64, messageId int64, text string, options ...MessageOption) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("sendMessage?chat_id=%v&reply_to_message_id=%v&text=%s", chatId, messageId, formatingText(text)))

	for _, opt := range options {
		err := opt(&url)
		if err != nil {
			log.Printf("error option %s", err)
		}
	}

	var result SendMessageResult
	err := t.Get(url.String(), &result)
	if err != nil {
		return result
	}

	return result
}

func (t *TelegramBot) EditMessage(chatId int64, messageId int64, text string, options ...MessageOption) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("editMessageText?chat_id=%v&message_id=%v&text=%s", chatId, messageId, formatingText(text)))

	for _, opt := range options {
		err := opt(&url)
		if err != nil {
			log.Printf("error option %s", err)
		}
	}

	var result SendMessageResult
	err := t.Get(url.String(), &result)
	if err != nil {
		return result
	}

	return result
}

func formatingText(text string) string {
	text = url.QueryEscape(text)

	return text
}

func (t *TelegramBot) SendChatAction(chatId int64, action ChatAction) Result {
	url := t.GetApiUrl(fmt.Sprintf("%s?chat_id=%v&action=%s", "sendChatAction", chatId, action))

	var result Result
	err := t.Get(url.String(), &result)
	if err != nil {
		return result
	}

	return result
}
