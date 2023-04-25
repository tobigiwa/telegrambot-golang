package bot

import (
	tele "gopkg.in/telebot.v3"
)

func MotivationFunc(c tele.Context) error {
	return c.Reply("_what your pick..._", MotivationInlineKeyboard(), tele.ModeMarkdown)

}

func GetTodaysQuote(c tele.Context) error {
	
	return c.Reply("_what your pick..._", MotivationInlineKeyboard(), tele.ModeMarkdown)

}
