package bot

import (
	"fmt"
	"sync"

	"github.com/tobigiwa/telegrambot-golang/internal/services"
	tele "gopkg.in/telebot.v3"
)

func getScheduledText(work func() ([]string, error)) string {
	msg, _ := TextResponse(work())
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
		go func(bot *tele.Bot, recepient int64) {
			defer wg.Done()
			bot.Send(&tele.User{ID: recepient}, getScheduledText(services.GetTodaysQuote), tele.ModeHTML)
		}(bot, user)
	}
	wg.Wait()
	a.Logger.WriteToStandarOutput(fmt.Sprintf("Success sent out %v scheduled messages", len(users)))
}

func (a Application) CronTodaysReligiousMessage(bot *tele.Bot) {
	users, err := a.Storage.AllIDs()
	if err != nil {
		a.Logger.LogError(err, "DB")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(users))

	for _, user := range users {
		go func(bot *tele.Bot, recepient int64) {
			defer wg.Done()

			bot.Send(&tele.User{ID: recepient}, getScheduledText(services.ScrapeBibleText), tele.ModeHTML)
			if au, err := resolveAudioMessgae(); err == nil {
				bot.Send(&tele.User{ID: recepient}, au)
			}

		}(bot, user)
	}
	wg.Wait()
	a.Logger.WriteToStandarOutput(fmt.Sprintf("Success sent out %v scheduled messages", len(users)))
}
