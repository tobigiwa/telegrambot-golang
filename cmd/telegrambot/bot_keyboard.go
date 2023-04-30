package main

import (
	tele "gopkg.in/telebot.v3"
)

var (
	startKeyboard = &tele.ReplyMarkup{ResizeKeyboard: false,
		OneTimeKeyboard: true,
		Placeholder:     `What would you like to do...`,
		// ForceReply:      true,
	}

	MotivationKeyboardBtn = startKeyboard.Text("Our Motivations 🧘‍♀️🧘‍♂️")
	ReligionKeyboardBtn   = startKeyboard.Text("Lord's Message ✝️️🛐")
	TherapyKeyboardBtn    = startKeyboard.Text("Meet the Therapist 💆🏾‍♀️💗")
	RemindernKeyboardBtn  = startKeyboard.Text("Reminder 🕰")
)

func StartKeyboard() *tele.ReplyMarkup {
	startKeyboard.Reply(
		startKeyboard.Row(MotivationKeyboardBtn, ReligionKeyboardBtn),
		startKeyboard.Row(TherapyKeyboardBtn),
		startKeyboard.Row(RemindernKeyboardBtn),
	)
	return startKeyboard
}

var (
	religionMessagesKeyboard = &tele.ReplyMarkup{ResizeKeyboard: false,
		OneTimeKeyboard: true,
		Placeholder:     `Blessed His the Word of the Lord...♱`,
	}

	BibleTextReligionMessageKeyboardBtn    = startKeyboard.Text("Today's Bible reading 🎚️")
	AudioReligionMessageKeyboardBtn        = startKeyboard.Text("Today's Audio message 🎵")
	AudioAndTextReligionMessageKeyboardBtn = startKeyboard.Text("Bible reading 🎚️ and Audio message🎵")
)

func ReligionMessageKeyboard() *tele.ReplyMarkup {
	religionMessagesKeyboard.Reply(
		religionMessagesKeyboard.Row(BibleTextReligionMessageKeyboardBtn, AudioReligionMessageKeyboardBtn),
		religionMessagesKeyboard.Row(AudioAndTextReligionMessageKeyboardBtn),
	)
	return religionMessagesKeyboard
}
