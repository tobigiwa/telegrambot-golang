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
	return fmt.Sprintf("%02d%02d", month, day)
}

func GetAudioMessage() error {
	var partAuidoURL string = "https://www.heartlight.org/audio/votd/2"

	monthAndDay := resolveAudioMessageDate()
	fullAudioURL := fmt.Sprintf("%v/%v.mp3", partAuidoURL, monthAndDay)

	res, err := http.Get(fullAudioURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	file, err := os.Create("assets/" + AudioFilename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func ScrapeBibleText() ([]string, error) {
	resp, err := http.Get("https://www.verseoftheday.com/#featured")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	message := strings.Split(doc.Find("#featured .scripture .bilingual-left").Text(), "—")[0]
	verse := doc.Find("#featured .scripture .bilingual-left a").Text()

	return []string{message, verse}, nil
}
