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

func (a Application) ScheduleTask(bot *tele.Bot, task func() ([]string, error), others ...interface{}) {
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

			bot.Send(&tele.User{ID: recepient}, getScheduledText(task), tele.ModeHTML)

			for _, fn := range others {
				if function, ok := fn.(func() (*tele.Audio, error)); ok {
					if au, err := function(); err == nil {
						bot.Send(&tele.User{ID: recepient}, au)
					}
				}
			}
		}(bot, user)
	}
	wg.Wait()
	a.Logger.WriteToStandarOutput(fmt.Sprintf("Success sent out %v scheduled messages", len(users)))
}

func DoNothing() {

}
