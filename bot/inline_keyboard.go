package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NumericInlineKeyboard = tgbot.NewInlineKeyboardMarkup(
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbot.NewInlineKeyboardButtonData("_222_", "222"),
		tgbot.NewInlineKeyboardButtonData("*333*", "333"),
	),
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonData("4", "4"),
		tgbot.NewInlineKeyboardButtonData("5", "5"),
		tgbot.NewInlineKeyboardButtonData("6", "6"),
	),
)
