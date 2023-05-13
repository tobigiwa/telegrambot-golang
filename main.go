package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tobigiwa/telegrambot-golang/internal/services"
	"github.com/tobigiwa/telegrambot-golang/internal/store"
	"github.com/tobigiwa/telegrambot-golang/logging"

	botBuild "github.com/tobigiwa/telegrambot-golang/bot"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// DATABSE
	conn, err := pgxpool.New(context.Background(), getDatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	db := store.NewRespository(conn)
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
	bot := botBuild.NewBot(getBotToken(), 10)

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
	app.Bot.Handle(&botBuild.GameKeyboardBtn, app.GameKeyboardHandleFunc)
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
	now := time.Now().UTC()
	lagos_time := now.In(time.FixedZone("WAT", 3600))
	s := gocron.NewScheduler(lagos_time.Location())

	s.Every(1).Day().At("6:30").Do(app.ScheduledTaskText, app.Bot, services.ScrapeBibleText)
	s.Every(1).Day().At("6:30").Do(app.ScheduledTaskMedia, app.Bot, botBuild.ResolveAudioMessge)

	s.Every(1).Day().At("7:30").Do(app.ScheduledTaskText, app.Bot, services.GetTodaysQuote)

	s.Every(1).Day().At("8:30").Do(app.ScheduledTaskMedia, app.Bot, botBuild.ResolveImageMessage)

	s.Every(1).Day().At("20:00").Do(app.ScheduledTaskText, app.Bot, services.GetRandomQuote)

	s.StartAsync()

	// polling from Telegram
	app.Bot.Start()

}

func getBotToken() (token string) {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok || token == "" {
		log.Fatal("No Token")
	}
	return token
}

func getDatabaseURL() (databaseURL string) {
	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok || databaseURL == "" {
		log.Fatal("No databaseURL")
	}
	return databaseURL
}
