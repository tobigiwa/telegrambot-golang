package main

import (
	"fmt"
	"time"

	"github.com/tobigiwa/telegrambot-golang/internal/services"

	tele "gopkg.in/telebot.v3"
)

// type BOT struct {
// 	Bot *tele.Bot
// }

func (a Application) StartHandlerFunc(c tele.Context) error {
	return c.Send(`<b><i>What would you like to do...</i></b>`, StartKeyboard(), tele.ModeHTML)

}

func (a Application) MotivationFunc(c tele.Context) error {
	return c.Reply("Good choice...", MotivationInlineKeyboard(), tele.ModeMarkdown)
}

func (a Application) GetTodaysQuoteFunc(c tele.Context) error {
	return c.Send(formatQuoteText(services.GetTodaysQuote), tele.ModeHTML)
}

func (a Application) GetRandomQuoteFunc(c tele.Context) error {
	return c.Send(formatQuoteText(services.GetRandomQuote), tele.ModeHTML)
}

func (a Application) GetRandomQuoteImageFunc(c tele.Context) error {
	res := services.GetRandomsQuoteImage()
	if res != nil {
		return c.Send(formatQuoteText(res), tele.ModeHTML)
	}
	p := &tele.Photo{File: tele.FromDisk("assets/image.jpeg")}
	return c.Send(p)
}

func (a Application) ReligionKeyboardHandlerFunc(c tele.Context) error {
	return c.Reply("***Blessed His the Word of the Lord...♱***", ReligionMessageKeyboard(), tele.ModeMarkdown)
}

func (a Application) GetBibleTextHandlerFunc(c tele.Context) error {
	return c.Reply(formatQuoteText(services.ScrapeBibleText), tele.ModeHTML)
}

func (a Application) GetAudioMessageHandlerFunc(c tele.Context) error {

	_, m, d := time.Now().Date()
	filename := fmt.Sprintf("Bible Reading for %v/%v", m, d)

	if checkIfFilePresent(services.AudioFilename) {
		a := &tele.Audio{File: tele.FromDisk("assets/" + services.AudioFilename), Title: filename, Performer: filename}
		return c.Reply(a)
	} else {
		res := services.GetAudioMessage()
		if res != nil {
			return c.Reply(formatQuoteText(res), tele.ModeHTML)
		}
		a := &tele.Audio{File: tele.FromDisk("assets/" + services.AudioFilename), Title: filename, Performer: filename}
		return c.Reply(a)
	}
}

func (a Application) GetBothAudioAndTextReligionMessageHandlerFunc(c tele.Context) error {
	a.GetBibleTextHandlerFunc(c)
	a.GetAudioMessageHandlerFunc(c)
	return nil
}
