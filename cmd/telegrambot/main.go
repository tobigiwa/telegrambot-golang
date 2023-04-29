package main

import (
	"log"
	"os"
	"time"

	build "github.com/tobigiwa/telegrambot-golang/bot"
	tele "gopkg.in/telebot.v3"
)

func getToken() string {
	BotToken, ok := os.LookupEnv("BOT_TOKEN1")
	if !ok || BotToken == "" {
		log.Fatal("No Token")
	}
	return BotToken
}

func main() {

	pref := tele.Settings{
		Token:  getToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot := build.NewBot(pref)
	log.Printf("Authorized on account... %s", bot.Me.Username)

	// keyboards
	bot.Handle(&build.MotivationKeyboardBtn, build.MotivationFunc)
	bot.Handle(&build.ReligionKeyboardBtn, build.ReligionKeyboardHandlerFunc)
	bot.Handle(&build.BibleTextReligionMessageKeyboardBtn, build.GetBibleTextHandlerFunc)
	bot.Handle(&build.AudioReligionMessageKeyboardBtn, build.GetAudioMessageHandlerFunc)
	bot.Handle(&build.AudioAndTextReligionMessageKeyboardBtn, build.GetBothAudioAndTextReligionMessageHandlerFunc)

	// inline keyboards
	bot.Handle(&build.GetTodaysQouteInlineKeyboardBtn, build.GetTodaysQuoteFunc)
	bot.Handle(&build.RandomQuotesKeyboardBtn, build.GetRandomQuoteFunc)
	bot.Handle(&build.ImageQoutesOnInlineKeyboardBtn, build.GetRandomQuoteImageFunc)

	// any text
	bot.Handle(tele.OnText, build.StartHandlerFunc, build.CheckMemberShip)

	build.InitializeAllCronTask(bot)

	bot.Start()

}
