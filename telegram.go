package telegram_bot

import (
	"encoding/json"
	"fmt"
	"io"
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

			for _, route := range t.Routes {
				go route(update, t)
			}

		}
		<-timeout.C
	}
}

func (t *TelegramBot) GetApiUrl(method string) string {
	url := fmt.Sprintf("%s%s/%s", t.ApiURL, t.Token, method)
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %s", err)
		return err
	}

	// Декодирование JSON-данных в структуру
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

func Text(f Handle) Handle {
	return func(update UpdateResult, bot *TelegramBot) {
		if !strings.HasPrefix(update.Message.Text, "/") {
			f(update, bot)
		}
	}
}

func Command(cmd string, f Handle) Handle {
	return func(update UpdateResult, bot *TelegramBot) {
		if strings.HasPrefix(update.Message.Text, fmt.Sprintf("/%s", cmd)) {
			f(update, bot)
		}
	}
}

func ReplyMarkup() ReplyKeyboardMarkup {
	return ReplyKeyboardMarkup{
		ResizeKeyboard: true,
	}
}

func (rm *ReplyKeyboardMarkup) Option() MessageOption {

	_bRm, _ := json.Marshal(rm)

	_res := make(map[string]interface{})

	json.Unmarshal(_bRm, &_res)
	res, _ := json.Marshal(_res)

	return MessageOption{
		Field: MOFReplyMarkup,
		Value: string(res),
	}
}

func (rm *ReplyKeyboardMarkup) Add(keyboard []KeyboardButton) {
	rm.Keyboard = append(rm.Keyboard, keyboard)
}

func (rm *ReplyKeyboardMarkup) Button(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

func (k *KeyboardButton) SetWebApp(url string) {
	k.WebApp = &WebAppInfo{Url: url}
}

func InlineMarkup() InlineKeyboardMarkup {
	return InlineKeyboardMarkup{}
}
func (i *InlineKeyboardMarkup) Option() MessageOption {

	cp := *i
	cp.InlineKeyboard = [][]InlineKeyboardButton{}

	for i, k := range i.InlineKeyboard {
		cp.InlineKeyboard = append(cp.InlineKeyboard, []InlineKeyboardButton{})
		for _, b := range k {
			if b.CallbackData != nil || b.Url != nil || b.SwitchInlineQuery != nil {
				cp.InlineKeyboard[i] = append(cp.InlineKeyboard[i], b)
			}
		}
	}

	_bRm, _ := json.Marshal(cp)

	_res := make(map[string]interface{})

	json.Unmarshal(_bRm, &_res)
	res, _ := json.Marshal(_res)

	return MessageOption{
		Field: MOFReplyMarkup,
		Value: string(res),
	}
}

func (i *InlineKeyboardMarkup) Add(keyboard []InlineKeyboardButton) {
	i.InlineKeyboard = append(i.InlineKeyboard, keyboard)
}

func (i *InlineKeyboardMarkup) Button(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{
		Text: text,
	}
}

func (ik *InlineKeyboardButton) SetUrl(url string) *InlineKeyboardButton {
	ik.Url = &url
	return ik
}

func (ik *InlineKeyboardButton) SetCallbackData(data string) *InlineKeyboardButton {
	ik.CallbackData = &data
	return ik
}

func (ik *InlineKeyboardButton) SetSwitchInlineQuery(data string) *InlineKeyboardButton {
	ik.SwitchInlineQuery = &data
	return ik
}

func RemoveKeyboard() ReplyKeyboardRemove {
	return ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
}

func (rkr *ReplyKeyboardRemove) Option() MessageOption {
	b, _ := json.Marshal(rkr)
	return MessageOption{
		Field: MOFReplyMarkup,
		Value: string(b),
	}
}

func NewForceReply() ForceReply {
	return ForceReply{ForceReply: true}
}

func (fr *ForceReply) Option() MessageOption {
	b, _ := json.Marshal(fr)
	return MessageOption{
		Field: MOFReplyMarkup,
		Value: string(b),
	}
}
func (fr *ForceReply) SetInputFieldPlaceholder(data string) {
	fr.InputFieldPlaceholder = &data
}
