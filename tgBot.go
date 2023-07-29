package telegram_bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Handle func(update UpdateResult, bot *TelegramBot)

type TelegramBot struct {
	ApiURL       string
	Token        string
	LastUpdateId int64
	Timeout      time.Duration
	Routes       []Handle
}

func NewBot(token string) (TelegramBot, error) {
	tg := TelegramBot{
		ApiURL:  "https://api.telegram.org/bot",
		Token:   token,
		Timeout: time.Second,
		Routes:  make([]Handle, 0),
	}

	if token == "" {
		return tg, fmt.Errorf("token is empty")
	}

	return tg, nil
}

func (t *TelegramBot) Polling() {
	timeout := time.NewTicker(t.Timeout)
	for {

		updates := t.GetUpdates(t.LastUpdateId+1, 0, 0)
		for _, update := range updates.Result {
			//обновляем счетчик
			if update.UpdateId > t.LastUpdateId {
				t.LastUpdateId = update.UpdateId
			}
			fmt.Println(update)
			for _, route := range t.Routes {
				go route(update, t)
			}

		}
		<-timeout.C
	}
}

func (t *TelegramBot) GetApiUrl(method string) string {
	url := fmt.Sprintf("%s%s/%s", t.ApiURL, t.Token, method)
	// println(url)
	return url
}

func (t *TelegramBot) Get(url string, resInterface interface{}) error {

	// Отправка GET-запроса
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %s", err)
		return err
	}
	defer response.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %s", err)
		return err
	}

	// Декодирование JSON-данных в структуру
	// fmt.Println(string(body))
	err = json.Unmarshal(body, resInterface)
	if err != nil {
		fmt.Printf("Ошибка при декодировании JSON: %s", err)
		return err
	}

	return nil
}

func (t *TelegramBot) AddHandle(f Handle) {
	t.Routes = append(t.Routes, f)
}

func Command(cmd string, f Handle) Handle {
	return func(update UpdateResult, bot *TelegramBot) {
		if strings.HasPrefix(update.Message.Text, fmt.Sprintf("/%s", cmd)) {
			f(update, bot)
		}
	}
}
