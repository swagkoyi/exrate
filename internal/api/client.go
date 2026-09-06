package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const currencies = "https://latest.currency-api.pages.dev/v1/currencies/%s.json"

func GetRate(input, target string) (float64, error) {
	input = strings.ToLower(strings.TrimSpace(input))
	target = strings.ToLower(strings.TrimSpace(target))

	rates, err := Fetch(input)
	if err != nil {
		return 0, err
	}

	rate, ok := rates[target]
	if !ok {
		return 0, errors.New("target not found")
	}

	return rate, nil
}

func GetRates(input string, target []string) (map[string]float64, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	rates, err := Fetch(input)
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)

	for _, tar := range target {
		tar = strings.ToLower(strings.TrimSpace(input))
		value, ok := rates[tar]
		if !ok {
			return nil, errors.New("tar not found")
		} else if ok {
			result[tar] = value
		}
	}
	return result, nil

}

func Fetch(input string) (map[string]float64, error) {
	url := fmt.Sprintf(currencies, input)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	/*firstUnmarshal*/
	raw := make(map[string]json.RawMessage)
	err = json.Unmarshal(body, &raw)
	if err != nil {
		log.Fatal(err)
	}

	value, ok := raw[input]
	if !ok {
		log.Fatal("wrong currency")
	}

	/*secondUnmarshal*/
	rates := make(map[string]float64)
	err = json.Unmarshal(value, &rates)
	if err != nil {
		log.Fatal(err)
	}

	return rates, nil
}
