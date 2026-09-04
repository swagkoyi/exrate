package client

import (
	"encoding/json"
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

	url := fmt.Sprintf(currencies, input)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	/*-------------------------------------------*/
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
	/*-------------------------------------------*/

	rate, ok := rates[target]
	if !ok {
		log.Fatal(0, "wrong zer0")
	}
	return rate, nil
}
