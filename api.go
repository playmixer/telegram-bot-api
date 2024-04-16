package telegram_bot

import (
	"fmt"
	"net/url"
)

func (t *TelegramBot) GetUpdates(offset, limit, timeout int64) Update {
	if limit == 0 {
		limit = 100
	}
	url := t.GetApiUrl(fmt.Sprintf("%s?offset=%v&limit=%v&timeout=%v", "getUpdates", offset, limit, timeout))
	// fmt.Println(url)

	var data Update
	err := t.Get(url, &data)
	if err != nil {
		return data
	}

	return data
}

func urlAddOption(url string, options ...MessageOption) string {
	for _, o := range options {
		switch o.Field {
		case MOFParseMode:
			url += fmt.Sprintf("&%s=%s", o.Field, o.Value)
		}
	}

	return url
}

func (t *TelegramBot) SendMessage(chatId int64, text string, options ...MessageOption) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("sendMessage?chat_id=%v&text=%s", chatId, formatingText(text)))
	url = urlAddOption(url, options...)

	var result SendMessageResult
	err := t.Get(url, &result)
	if err != nil {
		return result
	}

	return result
}

func (t *TelegramBot) ReplyToMessage(chatId int64, messageId int64, text string, options ...MessageOption) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("sendMessage?chat_id=%v&reply_to_message_id=%v&text=%s", chatId, messageId, formatingText(text)))
	url = urlAddOption(url, options...)

	var result SendMessageResult
	err := t.Get(url, &result)
	if err != nil {
		return result
	}

	return result
}

func (t *TelegramBot) EditMessage(chatId int64, messageId int64, text string, options ...MessageOption) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("editMessageText?chat_id=%v&message_id=%v&text=%s", chatId, messageId, formatingText(text)))
	url = urlAddOption(url, options...)

	var result SendMessageResult
	err := t.Get(url, &result)
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
	err := t.Get(url, &result)
	if err != nil {
		return result
	}

	return result
}
