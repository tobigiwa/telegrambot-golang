package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tobigiwa/telegrambot-golang/internal/store"
	"github.com/tobigiwa/telegrambot-golang/logging"

	botBuild "github.com/tobigiwa/telegrambot-golang/bot"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// DATABSE
	_, err := os.Create("db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	db := store.NewSQLiteRespository(conn)
	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// LOGGER
	logger, err := logging.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	// BOT
	bot := botBuild.NewBot(getToken(), 10)

	app := botBuild.Application{
		Bot:     bot,
		Storage: db,
		Logger:  logger,
	}

	// Other setups
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := cwd + "/assets"

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	app.Logger.WriteToStandarOutput(cwd)
	app.Logger.WriteToStandarOutput(fmt.Sprintf("Authorized on account... %s", app.Bot.Me.Username))

	// keyboards
	app.Bot.Handle(&botBuild.MotivationKeyboardBtn, app.MotivationKeyboardHandlerFunc)
	app.Bot.Handle(&botBuild.TherapyKeyboardBtn, app.TherapyKeyboardHandleFunc)
	app.Bot.Handle(&botBuild.RemindernKeyboardBtn, app.RemainderyKeyboardHandleFunc)
	app.Bot.Handle(&botBuild.ReligionKeyboardBtn, app.ReligionKeyboardHandlerFunc)
	app.Bot.Handle(&botBuild.BibleTextReligionMessageKeyboardBtn, app.GetBibleTextHandlerFunc)
	app.Bot.Handle(&botBuild.AudioReligionMessageKeyboardBtn, app.GetAudioMessageHandlerFunc)
	app.Bot.Handle(&botBuild.AudioAndTextReligionMessageKeyboardBtn, app.GetBothAudioAndTextReligionMessageHandlerFunc)
	app.Bot.Handle(&botBuild.BackToStartKeyboardBtn, app.BackToStartHanlerFunc)

	// inline keyboards
	app.Bot.Handle(&botBuild.GetTodaysQouteInlineKeyboardBtn, app.GetTodaysQuoteFunc)
	app.Bot.Handle(&botBuild.RandomQuotesInlineKeyboardBtn, app.GetRandomQuoteFunc)
	app.Bot.Handle(&botBuild.ImageQoutesOnInlineKeyboardBtn, app.GetRandomQuoteImageFunc)

	// any text
	app.Bot.Handle(tele.OnText, app.StartHandlerFunc, app.CheckMemberShip)

	// cron jobs
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("6:30").Do(app.CronTodaysQuote, app.Bot)

	s.StartAsync()

	// polling
	app.Bot.Start()

}

func getToken() (token string) {
	token, ok := os.LookupEnv("BOT_TOKEN1")
	if !ok || token == "" {
		log.Fatal("No Token")
	}
	return token
}
