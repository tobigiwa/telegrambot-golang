package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var AudioFilename string = fmt.Sprintf("audio_%v.mp3", resolveAudioMessageDate())

func resolveAudioMessageDate() string {
	_, month, day := time.Now().Date()
	return fmt.Sprintf("%02d%v", month, day)
}

func GetAudioMessage() []string {
	var partAuidoURL string = "https://www.heartlight.org/audio/votd/2"

	monthAndDay := resolveAudioMessageDate()
	fullAudioURL := fmt.Sprintf("%v/%v.mp3", partAuidoURL, monthAndDay)

	res, err := http.Get(fullAudioURL)
	if err != nil {
		return failedResponse
	}
	defer res.Body.Close()
	file, err := os.Create("assets/" + AudioFilename)
	if err != nil {
		return failedResponse
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return failedResponse
	}
	return nil
}

func ScrapeBibleText() []string {
	resp, err := http.Get("https://www.verseoftheday.com/#featured")
	if err != nil {
		return failedResponse
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return failedResponse
	}

	message := strings.Split(doc.Find("#featured .scripture .bilingual-left").Text(), "—")[0]
	verse := doc.Find("#featured .scripture .bilingual-left a").Text()

	return []string{message, verse}
}
