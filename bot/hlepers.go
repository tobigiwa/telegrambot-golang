package bot

import (
	"fmt"
	"log"
	"os"

	tele "gopkg.in/telebot.v3"
)

func NewBot(pref tele.Settings) *tele.Bot {
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func formatQuoteText[T FuncOrSlice](param T) string {

	switch v := any(param).(type) {
	case func() []string:
		res := v()
		return fmt.Sprintf("<b>%v</b>\n\n<i>%v</i>", res[0], res[1])
	case []string:
		return fmt.Sprintf("<b>%v</b>\n\n<i>%v</i>", v[0], v[1])
	default:
		fmt.Println("error type switch")
	}
	return ""
}

func checkIfFilePresent(filename string) bool {

	if _, err := os.Stat("assets/" + filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false

}
