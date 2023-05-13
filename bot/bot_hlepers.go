package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tobigiwa/telegrambot-golang/internal/services"
	tele "gopkg.in/telebot.v3"
)

var FailedRequest string = fmt.Sprintf("<b>%v</b>\n\n<i>%v</i>", "unable to fetch request", "please do try again")

func NewBot(token string, timeout int) *tele.Bot {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: time.Duration(timeout) * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func TextResponse(v []string, err error) (string, error) {
	if err != nil {
		return FailedRequest, err
	}
	if v == nil {
		return FailedRequest, nil
	}
	return fmt.Sprintf("<b>%v</b>\n\n<i>%v</i>", v[0], v[1]), nil

}

func checkIfFilePresent(filename string) bool {

	if _, err := os.Stat("assets/" + filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false

}

func ResolveAudioMessge() (*tele.Audio, error) {

	_, m, d := time.Now().Date()
	filename := fmt.Sprintf("Bible Reading for %v/%v", m, d)

	if checkIfFilePresent(services.AudioFilename) {
		a := &tele.Audio{File: tele.FromDisk("assets/" + services.AudioFilename), Title: filename, Performer: filename}
		return a, nil
	}
	err := services.GetAudioMessage()
	if err != nil {
		return nil, err
	}
	au := &tele.Audio{File: tele.FromDisk("assets/" + services.AudioFilename), Title: filename, Performer: filename}
	return au, nil

}

func ResolveImageMessage() (*tele.Photo, error) {

	err := services.GetRandomsQuoteImage()
	if err != nil {
		return nil, err
	}
	img := &tele.Photo{File: tele.FromDisk("assets/image.jpeg")}
	return img, nil

}
