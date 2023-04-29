package services

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var (
	todaysQuote      string   = "https://zenquotes.io/api/today"
	randomQuote      string   = "https://zenquotes.io/api/random"
	randomQuoteImage string   = "https://zenquotes.io/api/image"
	failedResponse   []string = []string{"unable to fetch request", "please do try again"}
)

func get(url string) ([]QuotesAndAuthors, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var a []QuotesAndAuthors
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func processQuote(list []QuotesAndAuthors) []string {
	return []string{list[0].Quote, list[0].Author}

}
func GetTodaysQuote() []string {
	a, err := get(todaysQuote)
	if err != nil {
		return failedResponse
	}
	return processQuote(a)
}

func GetRandomQuote() []string {
	a, err := get(randomQuote)
	if err != nil {
		return failedResponse
	}
	return processQuote(a)
}

func GetRandomsQuoteImage() []string {
	resp, err := http.Get(randomQuoteImage)
	if err != nil {
		return failedResponse
	}
	file, err := os.Create("assets/image.jpeg")
	if err != nil {
		return failedResponse
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return failedResponse
	}
	return nil
}
