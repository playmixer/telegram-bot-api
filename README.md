# telegram-bot-api

```
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	tg "github.com/playmixer/telegram-bot-api"
)

func start(update tg.UpdateResult, bot *tg.TelegramBot) {
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

func filter_function(substr string) tg.Handle {
	return func(update tg.UpdateResult, bot *tg.TelegramBot) {
		if strings.Contains(update.Message.Text, substr) {
			func(update tg.UpdateResult, bot *tg.TelegramBot) {
				msg := bot.SendMessage(update.Message.Chat.Id, fmt.Sprintf("%s contain %s", update.Message.Text, substr))
				if !msg.Ok {
					fmt.Println(msg.Description)
				}
			}(update, bot)
		}
	}
}

func main() {
	bot, err := tg.NewBot("<Telegram token>")
	if err != nil {
		log.Panic(err)
	}
	bot.AddHandle(tg.Command("start", start))
	bot.AddHandle(tg.Text(echo))
	bot.AddHandle(filter_function("123"))

	bot.Timeout = time.Second
	bot.Polling()
}

```
