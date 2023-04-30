package main

import (
	tele "gopkg.in/telebot.v3"
)

func (a Application) CheckMemberShip(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		UserID := c.Message().Chat.ID
		if user := a.Storage.IsUser(UserID); !user {
			go func(a Application) {
				username := c.Message().Chat.Username
				err := a.Storage.Insert(UserID, username)
				if err != nil {
					a.Logger.LogError(err, "DB")
				}
			}(a)
			c.Send(`***Welcome our dear new Friend***`, tele.ModeMarkdown)
		}
		return next(c)

	}
}
