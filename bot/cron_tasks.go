package bot

import (
	"fmt"
	"sync"

	tele "gopkg.in/telebot.v3"
)

func getScheduledText(work func() ([]string, error)) string {
	msg, _ := TextResponse(work())
	return msg
}

func (a Application) ScheduledTaskText(bot *tele.Bot, task func() ([]string, error)) {
	users, err := a.Storage.AllIDs()
	if err != nil {
		a.Logger.LogError(err, "DB")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(users))

	for _, user := range users {
		go func(bot *tele.Bot, receiver int64) {
			defer wg.Done()

			bot.Send(&tele.User{ID: receiver}, getScheduledText(task), tele.ModeHTML)
		}(bot, user)
	}
	wg.Wait()
	a.Logger.WriteToStandarOutput(fmt.Sprintf("Success sent out %v text scheduled messages", len(users)))
}

func (a Application) ScheduledTaskMedia(bot *tele.Bot, media ...interface{}) {
	users, err := a.Storage.AllIDs()
	if err != nil {
		a.Logger.LogError(err, "DB")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(users))

	for _, user := range users {
		go func(bot *tele.Bot, receiver int64) {
			defer wg.Done()

			for _, fn := range media {
				if function, ok := fn.(func() (*tele.Audio, error)); ok {
					if au, err := function(); err == nil {
						bot.Send(&tele.User{ID: receiver}, au)
					}
				}
				if function, ok := fn.(func() (*tele.Photo, error)); ok {
					if au, err := function(); err == nil {
						bot.Send(&tele.User{ID: receiver}, au)
					}
				}
			}
		}(bot, user)
	}
	wg.Wait()
	a.Logger.WriteToStandarOutput(fmt.Sprintf("Success sent out %v media scheduled messages", len(users)))
}
