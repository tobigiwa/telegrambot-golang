package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NumericKeyboard = NewReplyKeyboard(
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("🧘‍♀️"),
	),
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("⚕"),
	),

// tgbot.NewKeyboardButtonRow(
//
//	tgbot.NewKeyboardButton("🕰"),
//
// ),
)

func NewReplyKeyboard(rows ...[]tgbot.KeyboardButton) tgbot.ReplyKeyboardMarkup {
	var keyboard [][]tgbot.KeyboardButton

	keyboard = append(keyboard, rows...)

	return tgbot.ReplyKeyboardMarkup{
		ResizeKeyboard:        false,
		Keyboard:              keyboard,
		InputFieldPlaceholder: "hfhdd",
	}
}
