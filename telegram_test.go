package telegram_bot_test

import (
	"fmt"
	"testing"
	"time"

	tg "github.com/playmixer/telegram-bot-api/v3"
)

var (
	token        = "933196870:AAGHMTJUVFr3eeeUKgsCagiSNb0i3fBnSE0"
	chatId int64 = 432495963
)

func TestSendMessage(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	msg := bot.SendMessage(chatId, "Hello")
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}

func TestSendHTMLMessage(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	msg := bot.SendMessage(chatId, "<b>hello</b> <i>html</i> <u>code</u>", tg.StyleMarkdown(tg.MessageStyleHTML))
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}
func TestSendMarkdownMessage(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	options := []tg.MessageOption{
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2),
	}
	msg := bot.SendMessage(chatId, "text\n*bold*_italic_ __underline__ ~strike~", options...)
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}
func TestSendCodeMessage(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	options := []tg.MessageOption{
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2),
	}
	msg := bot.SendMessage(chatId, "```goleng\n<b>test</b>```", options...)
	if !msg.Ok {
		t.Fatal("cant send message")
	}

}

func TestEditMessage(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}
	options := []tg.MessageOption{
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2),
	}
	msg := bot.SendMessage(chatId, "```goleng\n<b>test</b>```", options...)
	if !msg.Ok {
		t.Fatal("cant send message")
	}
	time.Sleep(time.Second)
	msg = bot.EditMessage(msg.Result.Chat.Id, msg.Result.MessageId, "```goleng\n<i><b>test2</b></i>```",
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2))
}

func TestReplyKeyboardMarkup(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	rm := tg.ReplyMarkup()
	rm.ResizeKeyboard = true
	btn1 := rm.Button("тексе")
	btn1.SetWebApp("https://habr.com/ru/news/693250/")
	rm.Add([]tg.KeyboardButton{btn1})
	rm.Add([]tg.KeyboardButton{rm.Button("тексе2")})

	options := []tg.MessageOption{
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2),
		rm.Option(),
	}
	msg := bot.SendMessage(chatId, "```goleng\n<b>test</b>```", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}
}

func TestReplyInlineKeyboardMarkup(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	bot.AddHandle(tg.Text(func(update tg.UpdateResult, bot *tg.TelegramBot) {
		fmt.Println(update.Message.Text)
		fmt.Println(update.CallbackQuery)
	}))
	msg := bot.SendMessage(chatId, "text")

	im := tg.InlineMarkup()
	btn1 := *im.Button("текст").SetUrl("https://habr.com/ru/news/693250/")
	btn2 := *im.Button("```текст2 qwe\n qeqweq \n```").SetCallbackData("weq")
	// btn2.SetSwitchInlineQuery("weq")
	im.Add([]tg.InlineKeyboardButton{btn1, btn2})
	im.Add([]tg.InlineKeyboardButton{*im.Button("текст2").SetCallbackData("ewq")})

	options := []tg.MessageOption{
		im.Option(),
	}
	msg = bot.SendMessage(chatId, "text", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}

	bot.Polling()
}
func TestReplyKeyboardRemove(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	bot.AddHandle(tg.Text(func(update tg.UpdateResult, bot *tg.TelegramBot) {
		fmt.Println(update.Message.Text)
		fmt.Println(update.CallbackQuery)
	}))

	rm := tg.ReplyMarkup()
	rm.ResizeKeyboard = true
	btn1 := rm.Button("тексе")
	btn1.SetWebApp("https://habr.com/ru/news/693250/")
	rm.Add([]tg.KeyboardButton{btn1})
	rm.Add([]tg.KeyboardButton{rm.Button("тексе2")})

	options := []tg.MessageOption{
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2),
		rm.Option(),
	}
	msg := bot.SendMessage(chatId, "text", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}

	RemoveKeyboard := tg.RemoveKeyboard()
	msg = bot.SendMessage(chatId, "text", RemoveKeyboard.Option())
	if !msg.Ok {
		fmt.Println("cant send message", msg.Description)
	}

	bot.Polling()
}

func TestForceReply(t *testing.T) {
	bot, err := tg.NewBot(token)
	if err != nil {
		t.Fatal(err)
	}

	bot.AddHandle(tg.Text(func(update tg.UpdateResult, bot *tg.TelegramBot) {
		fmt.Println(update.Message.Text)
		fmt.Println(update.CallbackQuery)
	}))

	fr := tg.NewForceReply()
	fr.SetInputFieldPlaceholder("weq")

	options := []tg.MessageOption{
		tg.StyleMarkdown(tg.MessageStyleMarkdownV2),
		fr.Option(),
	}
	msg := bot.SendMessage(chatId, "text", options...)
	if !msg.Ok {
		t.Fatal("cant send message", msg.Description)
	}

	bot.Polling()
}
