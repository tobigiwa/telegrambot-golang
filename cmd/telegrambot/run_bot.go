package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	telegrambot "github.com/tobigiwa/telegrambot-golang/bot"
)

func main() {
	BotToken, ok := os.LookupEnv("BOT_TOKEN3")
	fmt.Println(BotToken)
	if !ok || BotToken == "" {
		log.Fatal("No Token")
	}

	fmt.Println(BotToken)
	bot := telegrambot.NewBot(BotToken)
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message != nil {
			chatID := update.Message.Chat.ID
			textMessage := update.Message.Text
			msg := tgbotapi.NewMessage(chatID, textMessage)

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					telegrambot.SetParseModeToMarkdownV2(&msg).Text = "😬 *Go \n _AWAY_*"
					goto send
				default:
					msg.Text = "COMMAND WAY"
					goto send
				}
			}

			switch textMessage {
			case "start":
				msg.ReplyToMessageID = update.Message.MessageID
				goto send

			default:
				msg.ReplyMarkup = telegrambot.NumericKeyboard
				goto send
			}

		send:
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}
}
