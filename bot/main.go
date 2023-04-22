package bot

import (
	"fmt"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot(BotToken string) *tgbot.BotAPI {
	bot, err := tgbot.NewBotAPI(BotToken)
	fmt.Println(BotToken)
	if err != nil {
		fmt.Println("occurred here")
		panic(err)
	}

	return bot
}

func SetParseModeToMarkdownV2(msg *tgbot.MessageConfig) *tgbot.MessageConfig {
	msg.ParseMode = "MarkdownV2"
	return msg
}
func SetParseModeToHtml(msg *tgbot.MessageConfig) *tgbot.MessageConfig {
	msg.ParseMode = "HTML"
	return msg
}