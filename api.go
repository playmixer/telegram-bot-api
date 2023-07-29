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

func (t *TelegramBot) SendMessage(chatId int64, text string) SendMessageResult {
	url := t.GetApiUrl(fmt.Sprintf("%s?chat_id=%v&text=%s", "sendMessage", chatId, formatingText(text)))
	println(url)
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
