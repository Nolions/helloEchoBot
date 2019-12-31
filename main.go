package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("<TELEGRAM_BOT_TOKEN>")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("<TELEGRAM_BOT_WEBHOOK_URL>" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServe("0.0.0.0:8080", nil)

	for update := range updates {

		log.Printf("%+v\n", update.Message)
		log.Printf("%+v\n", update.Message.Chat)
		log.Printf("%+v\n", update.Message.Text)

		if !update.Message.IsCommand() {
			c := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(c)
		}

	}
}
