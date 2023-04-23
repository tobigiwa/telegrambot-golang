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
	if !ok || BotToken == "" {
		log.Fatal("No Token")
	}

	bot := telegrambot.NewBot(BotToken)
	// bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

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
		}

		if update.CallbackQuery != nil {
			fmt.Print("\n\n\n\n\n\n\n\n\n")
			a := update.CallbackQuery.Data
			fmt.Println(a)
			//switch a {
			//case "todaysMotivation":
			//	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, a)
			//	if _, err := bot.Request(callback); err != nil {
			//		panic(err)
			//	}
			//	NewMsg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			//	goto send

			//}
		}

	send:
		if _, err := bot.Send(NewMsg); err != nil {
			log.Panic(err)
		}
	}
}
