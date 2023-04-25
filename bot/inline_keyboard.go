package bot

import (
	tele "gopkg.in/telebot.v3"
)

var (
	motivationInlineKeyboard = &tele.ReplyMarkup{ResizeKeyboard: true}

	GetTodaysQouteInlineKeyboardBtn = motivationInlineKeyboard.Data("See todays Quote", "seeQuote", "seeQuote")
	QoutesOnInlineKeyboardBtn       = motivationInlineKeyboard.Data("Get quote on love, peace, money", "QuoteOn", "QuoteOn")
	RandomQuotesKeyboardBtn         = motivationInlineKeyboard.Data("See any quote", "anyQuote", "anyQuote")
)

func MotivationInlineKeyboard() *tele.ReplyMarkup {
	motivationInlineKeyboard.Inline(
		motivationInlineKeyboard.Row(GetTodaysQouteInlineKeyboardBtn),
		motivationInlineKeyboard.Row(QoutesOnInlineKeyboardBtn),
		motivationInlineKeyboard.Row(RandomQuotesKeyboardBtn),
	)
	return motivationInlineKeyboard
}
