package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/tobigiwa/telegrambot-golang/internal/store"
	"github.com/tobigiwa/telegrambot-golang/logging"

	tele "gopkg.in/telebot.v3"
)

type Database interface {
	// IsUser returns true if user is found in the db and false otherwise.
	IsUser(int64) bool
	Delete(int64) error
	Insert(int64, string) error
	All() ([]store.USER, error)
}

type Logger interface {
	LogInfo(string, string)
	LogError(error, string)
	LogFatal(error, string)
	WriteToStandarOutput(string)
}

// Application is the monolothic struct for the application
type Application struct {
	// Bot holds the Bot instance
	Bot *tele.Bot
	// Storage holds the database instance
	Storage Database
	//Logger holds the logger instance
	Logger Logger
}

func main() {
	// DATABSE
	conn, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	db := store.NewSQLiteRespository(conn)

	// LOGGER
	logger, err := logging.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	// BOT
	bot := NewBot(getToken(), 10)

	app := Application{
		Bot:     bot,
		Storage: db,
		Logger:  logger,
	}
	app.Logger.WriteToStandarOutput(fmt.Sprintf("Authorized on account... %s", app.Bot.Me.Username))

	// keyboards
	app.Bot.Handle(&MotivationKeyboardBtn, app.MotivationFunc)
	app.Bot.Handle(&ReligionKeyboardBtn, app.ReligionKeyboardHandlerFunc)
	app.Bot.Handle(&BibleTextReligionMessageKeyboardBtn, app.GetBibleTextHandlerFunc)
	app.Bot.Handle(&AudioReligionMessageKeyboardBtn, app.GetAudioMessageHandlerFunc)
	app.Bot.Handle(&AudioAndTextReligionMessageKeyboardBtn, app.GetBothAudioAndTextReligionMessageHandlerFunc)

	// inline keyboards
	app.Bot.Handle(&GetTodaysQouteInlineKeyboardBtn, app.GetTodaysQuoteFunc)
	app.Bot.Handle(&RandomQuotesInlineKeyboardBtn, app.GetRandomQuoteFunc)
	app.Bot.Handle(&ImageQoutesOnInlineKeyboardBtn, app.GetRandomQuoteImageFunc)

	// any text
	app.Bot.Handle(tele.OnText, app.StartHandlerFunc, app.CheckMemberShip)

	// cron jobs
	InitializeAllCronTask(app.Bot)

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
