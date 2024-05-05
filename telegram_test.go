package telegram_bot_test

import (
	"fmt"
	"testing"
	"time"

	telegram_bot "github.com/playmixer/telegram-bot-api/v2"
)

var (
	token        = "933196870:AAGHMTJUVFr3eeeUKgsCagiSNb0i3fBnSE0"
	chatId int64 = 432495963
)

func TestSendMessage(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	msg := bot.SendMessage(chatId, "Hello")
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}

func TestSendHTMLMessage(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	msg := bot.SendMessage(chatId, "<b>hello</b> <i>html</i> <u>code</u>", telegram_bot.MessageOption{Field: telegram_bot.MOFParseMode, Value: "HTML"})
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}
func TestSendMarkdownMessage(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	options := []telegram_bot.MessageOption{
		{Field: telegram_bot.MOFParseMode, Value: string(telegram_bot.MessageStyleMarkdownV2)},
	}
	msg := bot.SendMessage(chatId, "text\n*bold*_italic_ __underline__ ~strike~", options...)
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}
func TestSendCodeMessage(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	options := []telegram_bot.MessageOption{
		{Field: telegram_bot.MOFParseMode, Value: string(telegram_bot.MessageStyleMarkdownV2)},
	}
	msg := bot.SendMessage(chatId, "```goleng\n<b>test</b>```", options...)
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}

func TestEditMessage(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	options := []telegram_bot.MessageOption{
		{Field: telegram_bot.MOFParseMode, Value: string(telegram_bot.MessageStyleMarkdownV2)},
	}
	msg := bot.SendMessage(chatId, "```goleng\n<b>test</b>```", options...)
	if !msg.Ok {
		t.Fatal("cant send message")
	}
	time.Sleep(time.Second)
	msg = bot.EditMessage(msg.Result.Chat.Id, msg.Result.MessageId, "```goleng\n<i><b>test2</b></i>```", telegram_bot.MessageOption{Field: telegram_bot.MOFParseMode, Value: "MarkdownV2"})
}

func TestReplyKeyboardMarkup(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	rm := telegram_bot.ReplyMarkup()
	rm.ResizeKeyboard = true
	btn1 := rm.Button("тексе")
	btn1.SetWebApp("https://habr.com/ru/news/693250/")
	rm.Add([]telegram_bot.KeyboardButton{btn1})
	rm.Add([]telegram_bot.KeyboardButton{rm.Button("тексе2")})

	options := []telegram_bot.MessageOption{
		{Field: telegram_bot.MOFParseMode, Value: string(telegram_bot.MessageStyleMarkdownV2)},
		rm.Option(),
	}
	msg := bot.SendMessage(chatId, "```goleng\n<b>test</b>```", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}
}

func TestReplyInlineKeyboardMarkup(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	bot.AddHandle(telegram_bot.Text(func(update telegram_bot.UpdateResult, bot *telegram_bot.TelegramBot) {
		fmt.Println(update.Message.Text)
		fmt.Println(update.CallbackQuery)
	}))
	msg := bot.SendMessage(chatId, "text")

	im := telegram_bot.InlineMarkup()
	btn1 := *im.Button("текст").SetUrl("https://habr.com/ru/news/693250/")
	btn2 := *im.Button("```текст2 qwe\n qeqweq \n```").SetCallbackData("weq")
	// btn2.SetSwitchInlineQuery("weq")
	im.Add([]telegram_bot.InlineKeyboardButton{btn1, btn2})
	im.Add([]telegram_bot.InlineKeyboardButton{*im.Button("текст2").SetCallbackData("ewq")})

	options := []telegram_bot.MessageOption{
		im.Option(),
	}
	msg = bot.SendMessage(chatId, "text", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}

	bot.Polling()
}
func TestReplyKeyboardRemove(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	bot.AddHandle(telegram_bot.Text(func(update telegram_bot.UpdateResult, bot *telegram_bot.TelegramBot) {
		fmt.Println(update.Message.Text)
		fmt.Println(update.CallbackQuery)
	}))

	rm := telegram_bot.ReplyMarkup()
	rm.ResizeKeyboard = true
	btn1 := rm.Button("тексе")
	btn1.SetWebApp("https://habr.com/ru/news/693250/")
	rm.Add([]telegram_bot.KeyboardButton{btn1})
	rm.Add([]telegram_bot.KeyboardButton{rm.Button("тексе2")})

	options := []telegram_bot.MessageOption{
		{Field: telegram_bot.MOFParseMode, Value: string(telegram_bot.MessageStyleMarkdownV2)},
		rm.Option(),
	}
	msg := bot.SendMessage(chatId, "text", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}

	RemoveKeyboard := telegram_bot.RemoveKeyboard()
	msg = bot.SendMessage(chatId, "text", RemoveKeyboard.Option())
	if !msg.Ok {
		fmt.Println("cant send message", msg.Description)
	}

	bot.Polling()
}

func TestForceReply(t *testing.T) {
	bot, err := telegram_bot.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	bot.AddHandle(telegram_bot.Text(func(update telegram_bot.UpdateResult, bot *telegram_bot.TelegramBot) {
		fmt.Println(update.Message.Text)
		fmt.Println(update.CallbackQuery)
	}))

	fr := telegram_bot.NewForceReply()
	fr.SetInputFieldPlaceholder("weq")

	options := []telegram_bot.MessageOption{
		{Field: telegram_bot.MOFParseMode, Value: string(telegram_bot.MessageStyleMarkdownV2)},
		fr.Option(),
	}
	msg := bot.SendMessage(chatId, "text", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}

	bot.Polling()
}
