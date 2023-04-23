package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NumericKeyboard = customNewReplyKeyboard(customReplyKeyboardMarkup{ResizeKeyboard: false, InputFieldPlaceholder: "What would you lke to do now..."},
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("Our Motivations 🧘‍♀️"),
		tgbot.NewKeyboardButton("Remainder system 🕰"),
	),
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("Therapy 💆💚"),
	),
)

type customReplyKeyboardMarkup struct {
	ResizeKeyboard        bool
	OneTimeKeyboard       bool
	Selective             bool
	InputFieldPlaceholder string
}

func customNewReplyKeyboard(configs customReplyKeyboardMarkup, rows ...[]tgbot.KeyboardButton) tgbot.ReplyKeyboardMarkup {
	var keyboard [][]tgbot.KeyboardButton

	keyboard = append(keyboard, rows...)

	return tgbot.ReplyKeyboardMarkup{
		ResizeKeyboard:        configs.ResizeKeyboard,
		Keyboard:              keyboard,
		InputFieldPlaceholder: configs.InputFieldPlaceholder,
		Selective:             configs.Selective,
	}
}
