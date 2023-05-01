package services

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var (
	todaysQuote      string = "https://zenquotes.io/api/today"
	randomQuote      string = "https://zenquotes.io/api/random"
	randomQuoteImage string = "https://zenquotes.io/api/image"
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

func processQuote(list []QuotesAndAuthors) ([]string, error) {
	return []string{list[0].Quote, list[0].Author}, nil

}
func GetTodaysQuote() ([]string, error) {
	a, err := get(todaysQuote)
	if err != nil {
		return nil, err
	}
	return processQuote(a)
}

func GetRandomQuote() ([]string, error) {
	a, err := get(randomQuote)
	if err != nil {
		return nil, err
	}
	return processQuote(a)
}

func GetRandomsQuoteImage() error {
	resp, err := http.Get(randomQuoteImage)
	if err != nil {
		return err
	}
	file, err := os.Create("assets/image.jpeg")
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
