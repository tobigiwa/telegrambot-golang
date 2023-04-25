package main

import (
	"log"
	"os"
	"time"

	build "github.com/tobigiwa/telegrambot-golang/bot"
	tele "gopkg.in/telebot.v3"
)

func getToken() string {
	BotToken, ok := os.LookupEnv("BOT_TOKEN3")
	if !ok || BotToken == "" {
		log.Fatal("No Token")
	}
	return BotToken
}

func main() {
	pref := tele.Settings{
		Token:  getToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		// Verbose:   true,
	}

	bot := build.NewBot(pref)
	log.Printf("Authorized on account %s", bot.Me.Username)

	bot.Handle(&build.MotivationKeyboardBtn, build.MotivationFunc)

	bot.Handle(&build.GetTodaysQouteInlineKeyboardBtn, )

	bot.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(`<b><i>What would you like to do...</i></b>`, build.StartKeyboard(), tele.ModeHTML)
	})

	bot.Start()
}
