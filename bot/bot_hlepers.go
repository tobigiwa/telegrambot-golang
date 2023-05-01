package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

type FuncOrSlice interface {
	func() []string | []string
}

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
		return fmt.Sprintf("<b>%v</b>\n\n<i>%v</i>", "unable to fetch request", "please do try again"), err
	}
	if v == nil {
		return fmt.Sprintf("<b>%v</b>\n\n<i>%v</i>", "unable to fetch request", "please do try again"), nil
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
