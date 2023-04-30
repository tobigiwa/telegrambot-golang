package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	quote "github.com/tobigiwa/telegrambot-golang/internal/services"
	tele "gopkg.in/telebot.v3"
)

func GetScheduleTodayQoute() string {
	return formatQuoteText(quote.GetTodaysQuote)
}

func InitializeAllCronTask(bot *tele.Bot) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("6:30").Do(func() {
		_, err := bot.Send(&tele.User{ID: int64(987854119)}, GetScheduleTodayQoute(), tele.ModeHTML)
		if err != nil {
			fmt.Println(err)
		}
	})

	s.StartAsync()
}
