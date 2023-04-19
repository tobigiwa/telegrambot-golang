package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbot.NewReplyKeyboard(
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("1"),
		tgbot.NewKeyboardButton("2"),
		tgbot.NewKeyboardButton("3"),
	),
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("4"),
		tgbot.NewKeyboardButton("5"),
		tgbot.NewKeyboardButton("6"),
	),
)

var numericInlineKeyboard = tgbot.NewInlineKeyboardMarkup(
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbot.NewInlineKeyboardButtonData("222", "222"),
		tgbot.NewInlineKeyboardButtonData("333", "333"),
	),
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonData("4", "4"),
		tgbot.NewInlineKeyboardButtonData("5", "5"),
		tgbot.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	BotToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok || BotToken == "" {
		log.Fatal("No Token")
	}
	bot, err := tgbot.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	fmt.Printf("%+v\n\n%T\n\n%v\n\n\n", bot, bot, prettyPrint(bot))

	updateConfig := tgbot.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message != nil {
			msg := tgbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericInlineKeyboard
			case "close":
				msg.ReplyMarkup = tgbot.NewRemoveKeyboard(true)
			}
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbot.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg := tgbot.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

	}

}
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
