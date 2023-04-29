package bot

import (
	tele "gopkg.in/telebot.v3"
)

var (
	motivationInlineKeyboard = &tele.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		ForceReply:      true,
		Placeholder:     "click on any of the buttons...",
	}

	GetTodaysQouteInlineKeyboardBtn = motivationInlineKeyboard.Data("Todays Quote", "seeQuote", "seeQuote")
	RandomQuotesKeyboardBtn         = motivationInlineKeyboard.Data("Any quote", "anyQuote", "anyQuote")
	ImageQoutesOnInlineKeyboardBtn  = motivationInlineKeyboard.Data("Motivational Image", "QuoteOn", "image")
)

func MotivationInlineKeyboard() *tele.ReplyMarkup {
	motivationInlineKeyboard.Inline(
		motivationInlineKeyboard.Row(GetTodaysQouteInlineKeyboardBtn, RandomQuotesKeyboardBtn),
		motivationInlineKeyboard.Row(ImageQoutesOnInlineKeyboardBtn),
	)
	return motivationInlineKeyboard
}
