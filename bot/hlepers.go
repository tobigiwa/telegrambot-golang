package bot

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

func NewBot(pref tele.Settings) *tele.Bot {
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
