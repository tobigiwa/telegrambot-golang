package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func CheckMemberShip(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		chatID := c.Message().Chat.ID
		fmt.Print("\n\n\n\n", chatID, "\n\n\n\n")
		return next(c)
	}
}
