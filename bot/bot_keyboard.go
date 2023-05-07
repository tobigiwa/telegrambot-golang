package bot

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
	GameKeyboardBtn       = startKeyboard.Text("Games 🎮	🧩")
)

func StartKeyboard() *tele.ReplyMarkup {
	startKeyboard.Reply(
		startKeyboard.Row(MotivationKeyboardBtn, ReligionKeyboardBtn),
		startKeyboard.Row(TherapyKeyboardBtn),
		startKeyboard.Row(RemindernKeyboardBtn, GameKeyboardBtn),
	)
	return startKeyboard
}

var (
	religionMessagesKeyboard = &tele.ReplyMarkup{ResizeKeyboard: false,
		OneTimeKeyboard: true,
		Placeholder:     `Blessed His the Word of the Lord...♱`,
	}

	BibleTextReligionMessageKeyboardBtn    = religionMessagesKeyboard.Text("Today's Bible reading 🎚️")
	AudioReligionMessageKeyboardBtn        = religionMessagesKeyboard.Text("Today's Audio message 🎵")
	AudioAndTextReligionMessageKeyboardBtn = religionMessagesKeyboard.Text("Bible reading 🎚️ and Audio message🎵")
	BackToStartKeyboardBtn                 = religionMessagesKeyboard.Text("⬅⬅ Back to Start")
)

func ReligionMessageKeyboard() *tele.ReplyMarkup {
	religionMessagesKeyboard.Reply(
		religionMessagesKeyboard.Row(BibleTextReligionMessageKeyboardBtn, AudioReligionMessageKeyboardBtn),
		religionMessagesKeyboard.Row(AudioAndTextReligionMessageKeyboardBtn),
		religionMessagesKeyboard.Row(BackToStartKeyboardBtn),
	)
	return religionMessagesKeyboard
}
