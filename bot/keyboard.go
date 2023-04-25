package bot

import (
	tele "gopkg.in/telebot.v3"
)

var (
	startKeyboard = &tele.ReplyMarkup{ResizeKeyboard: false}

	MotivationKeyboardBtn = startKeyboard.Text("Our Motivations 🧘‍♀️🧘‍♂️")
	TherapyKeyboardBtn    = startKeyboard.Text("Meet the Therapist 💆🏾‍♀️💗")
	RemindernKeyboardBtn  = startKeyboard.Text("Reminder 🕰")
)

func StartKeyboard() *tele.ReplyMarkup {
	startKeyboard.Reply(
		startKeyboard.Row(MotivationKeyboardBtn),
		startKeyboard.Row(TherapyKeyboardBtn),
		startKeyboard.Row(RemindernKeyboardBtn),
	)
	return startKeyboard
}
