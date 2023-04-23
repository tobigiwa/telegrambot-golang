package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var FromBaseKeyboardInlineKeyboard = tgbot.NewInlineKeyboardMarkup(
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonData("Show today's motivation", "todaysMotivation"),
	),
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonData("Random motivation", "randomMotivation"),
	),
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonData("Motivation on...", "motivationOn"),
	),
)
