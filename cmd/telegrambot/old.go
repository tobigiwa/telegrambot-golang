package main

// import (
// 	"log"
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	telegrambot "github.com/tobigiwa/telegrambot-golang/bot"
// )

// var bot = telegrambot.NewBot(getToken())

// func main() {

// 	updateConfig := tgbotapi.NewUpdate(0)
// 	updateConfig.Timeout = 60
// 	updates := bot.GetUpdatesChan(updateConfig)

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	for update := range updates {

// 		chatID := update.Message.Chat.ID
// 		NewMsg := tgbotapi.NewMessage(chatID, update.Message.Text)
// 		if update.Message.IsCommand() {

// 			switch update.Message.Command() {
// 			case "start":
// 				telegrambot.SetParseModeToMarkdownV2(&NewMsg).Text = "😬 *Go \n _AWAY_*"
// 				SendNewMessage(NewMsg)
// 			default:
// 				NewMsg.Text = "COMMAND WAY"
// 				SendNewMessage(NewMsg)
// 			}
// 		}

// 		if update.Message != nil && update.Message.Text != "" {

// 			switch update.Message.Text {
// 			case "Our Motivations 🧘‍♀️":
// 				NewMsg.ReplyMarkup = telegrambot.FromBaseKeyboardInlineKeyboard
// 				NewMsg.ReplyToMessageID = update.Message.MessageID
// 				SendNewMessage(NewMsg)
// 			default:
// 				NewMsg.ReplyMarkup = telegrambot.NumericKeyboard
// 				SendNewMessage(NewMsg)
// 			}
// 		}

// 		if update.CallbackQuery != nil {

// 			callback := tgbotapi.NewCallback(update.CallbackQuery.ID,
// 				update.CallbackQuery.Data)
// 			if _, err := bot.Request(callback); err != nil {
// 				panic(err)
// 			}
// 			NewMsg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
// 			SendNewMessage(NewMsg)
// 		}

// 	}
// }

// func SendNewMessage(msg tgbotapi.Chattable) {
// 	if _, err := bot.Send(msg); err != nil {
// 		log.Panic(err)
// 	}
// }

// func getToken() string {
// 	BotToken, ok := os.LookupEnv("BOT_TOKEN2")
// 	if !ok || BotToken == "" {
// 		log.Fatal("No Token")
// 	}
// 	return BotToken
// }

// func SetParseModeToMarkdownV2(msg *tgbot.MessageConfig) *tgbot.MessageConfig {
// 	msg.ParseMode = "MarkdownV2"
// 	return msg
// }
// func SetParseModeToHtml(msg *tgbot.MessageConfig) *tgbot.MessageConfig {
// 	msg.ParseMode = "HTML"
// 	return msg
// }
