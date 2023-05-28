package bot

import (
	"sync"

	tele "gopkg.in/telebot.v3"
)

func (a Application) CheckMemberShip(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		UserID := c.Message().Chat.ID

		var wg sync.WaitGroup
		wg.Add(1)
		if user := a.Storage.IsUser(UserID); !user {

			go func(a Application) {
				defer wg.Done()
				username := c.Message().Chat.Username
				err := a.Storage.Insert(UserID, username)
				if err != nil {
					a.Logger.LogError(err, "DB")
					return
				}
			}(a)

			c.Send(`***Welcome our dear new Friend, we noticed it your first time here...***`, tele.ModeMarkdown)
			wg.Wait()
		}
		return next(c)

	}
}
