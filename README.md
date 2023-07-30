# telegram-bot-api

```
package main

import (
	"fmt"
	"log"
	"time"

	tg "github.com/playmixer/telegram-bot-api"
)

func start(update tg.UpdateResult, bot *tg.TelegramBot) {
	// fmt.Println(update.Message.Text)
	msg := bot.SendMessage(update.Message.Chat.Id, "Старт")
	if !msg.Ok {
		fmt.Println(msg.Description)
	}
}

func echo(update tg.UpdateResult, bot *tg.TelegramBot) {
	msg := bot.ReplyToMessage(update.Message.Chat.Id, update.Message.MessageId, update.Message.Text)
	if !msg.Ok {
		fmt.Println(msg.Description)
	}  

	bot.SendChatAction(msg.Result.Chat.Id, tg.TYPING)

	time.Sleep(time.Duration(time.Second * 2))
	msg = bot.EditMessage(msg.Result.Chat.Id, msg.Result.MessageId, "edited")
	if !msg.Ok {
		fmt.Println(msg.Description)
	}
}

func main() {
	bot, err := tg.NewBot("<TELEGRAM TOKEN>")
	if err != nil {
		log.Panic(err)
	}
	bot.AddHandle(tg.Command("start", start))
	bot.AddHandle(echo)

	bot.Timeout = time.Second
	bot.Polling()
}
```
