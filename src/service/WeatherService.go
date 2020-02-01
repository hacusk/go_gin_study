package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"example.com/go_gin_study/src/model"
)

// GetWeather return weather info
func GetWeather() {

	values := url.Values{}
	values.Add("city", "016010")
	resp, err := http.Get("http://weather.livedoor.com/forecast/webservice/json/v1" + "?" + values.Encode())

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	ParseWeather(resp)
}

// ParseWeather http.Response to Parse Json
func ParseWeather(resp *http.Response) {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	data := new(model.Weather)

	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	fmt.Println(data)
}
