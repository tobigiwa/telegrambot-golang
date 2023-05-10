package bot

import (
	"github.com/tobigiwa/telegrambot-golang/internal/services"

	tele "gopkg.in/telebot.v3"
)

func (a Application) StartHandlerFunc(c tele.Context) error {
	return c.Send(`<b><i>What would you like to do...</i></b>`, StartKeyboard(), tele.ModeHTML)
}

func (a Application) MotivationKeyboardHandlerFunc(c tele.Context) error {
	return c.Reply("Good choice...", MotivationInlineKeyboard(), tele.ModeMarkdown)
}

func (a Application) TherapyKeyboardHandleFunc(c tele.Context) error {
	return c.Reply("***Not implemented***", tele.ModeMarkdown)
}
func (a Application) RemainderyKeyboardHandleFunc(c tele.Context) error {
	return c.Reply("***Not implemented***", tele.ModeMarkdown)
}

func (a Application) GameKeyboardHandleFunc(c tele.Context) error {
	return c.Reply("***Not implemented***", tele.ModeMarkdown)
}

func (a Application) GetTodaysQuoteFunc(c tele.Context) error {
	msg, err := TextResponse(services.GetTodaysQuote())
	if err != nil {
		a.Logger.LogError(err, "SERVICES")
	}
	return c.Send(msg, tele.ModeHTML)
}

func (a Application) GetRandomQuoteFunc(c tele.Context) error {
	msg, err := TextResponse(services.GetRandomQuote())
	if err != nil {
		a.Logger.LogError(err, "SERVICES")
	}
	return c.Send(msg, tele.ModeHTML)
}

func (a Application) GetRandomQuoteImageFunc(c tele.Context) error {
	msg, err := TextResponse(nil, services.GetRandomsQuoteImage())
	if err != nil {
		a.Logger.LogError(err, "SERVICES")
		return c.Send(msg, tele.ModeHTML)
	}
	p := &tele.Photo{File: tele.FromDisk("assets/image.jpeg")}
	return c.Send(p)
}

func (a Application) ReligionKeyboardHandlerFunc(c tele.Context) error {
	return c.Reply("***Blessed His the Word of the Lord...♱***", ReligionMessageKeyboard(), tele.ModeMarkdown)
}

func (a Application) BackToStartHanlerFunc(c tele.Context) error {
	return a.StartHandlerFunc(c)
}

func (a Application) GetBibleTextHandlerFunc(c tele.Context) error {
	msg, err := TextResponse(services.ScrapeBibleText())
	if err != nil {
		a.Logger.LogError(err, "SERVICES")
	}
	return c.Reply(msg, tele.ModeHTML)
}

func (a Application) GetAudioMessageHandlerFunc(c tele.Context) error {
	audioFile, err := ResolveAudioMessgae()
	if err != nil {
		return c.Reply(FailedRequest, tele.ModeHTML)
	}
	return c.Reply(audioFile)
}

func (a Application) GetBothAudioAndTextReligionMessageHandlerFunc(c tele.Context) error {
	a.GetBibleTextHandlerFunc(c)
	a.GetAudioMessageHandlerFunc(c)
	return nil
}
