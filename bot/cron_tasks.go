package bot

import (
	"fmt"
	"sync"

	"github.com/tobigiwa/telegrambot-golang/internal/services"
	tele "gopkg.in/telebot.v3"
)

func (a Application) getScheduleTodayQoute() string {
	msg, err := TextResponse(services.GetTodaysQuote())
	if err != nil {
		a.Logger.LogError(err, "SERVICES")
	}
	return msg
}

func (a Application) CronTodaysQuote(bot *tele.Bot) {
	users, err := a.Storage.AllIDs()
	if err != nil {
		a.Logger.LogError(err, "DB")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(users))

	for _, user := range users {
		go func(b *tele.Bot, recepient int64) {
			defer wg.Done()
			bot.Send(&tele.User{ID: recepient}, a.getScheduleTodayQoute(), tele.ModeHTML)
		}(bot, user)
	}
	wg.Wait()
	a.Logger.WriteToStandarOutput(fmt.Sprintf("Success sent out %v scheduled messages", len(users)))
}
