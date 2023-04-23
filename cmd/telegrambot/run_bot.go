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

		if update.Message == nil { // if no message is sent
			continue
		}

		chatID := update.Message.Chat.ID
		NewMsg := tgbotapi.NewMessage(chatID, "...")

		if update.Message != nil {
			textMessage := update.Message.Text // handles all `text` messages
			if textMessage != "" {
				switch textMessage {
				case "Our Motivations 🧘‍♀️":
					NewMsg.ReplyMarkup = telegrambot.FromBaseKeyboardInlineKeyboard
					NewMsg.ReplyToMessageID = update.Message.MessageID
					goto send

				default:
					NewMsg.ReplyMarkup = telegrambot.NumericKeyboard
					goto send
				}
			}

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					telegrambot.SetParseModeToMarkdownV2(&NewMsg).Text = "😬 *Go \n _AWAY_*"
					goto send
				default:
					NewMsg.Text = "COMMAND WAY"
					goto send
				}
			}
		} else if update.CallbackQuery != nil {
			fmt.Print("\n\n\n\n\n\n\n\n\n")
			a := update.CallbackQuery.Data
			fmt.Println(a)
			switch a {
			case "todaysMotivation":
				telegrambot.SetParseModeToMarkdownV2(&NewMsg).Text = "*goooood*"
				goto send

			}
		}

	send:
		if _, err := bot.Send(NewMsg); err != nil {
			log.Panic(err)
		}
	}
}
