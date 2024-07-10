package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	requestUrl := "http://api.weatherapi.com/v1/forecast.json"
	apiKey := "c26061f634ea4804890113753241007"
	city := "lagos"
	days := "3"

	res, err := http.Get(fmt.Sprintf("%s?key=%s&q=%s&days=%s&alerts=no&aqi=no", requestUrl, apiKey, city, days))

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	if res.StatusCode != 200 {
		panic("Weather api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
