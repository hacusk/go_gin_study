package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"example.com/go_gin_study/src/model"
)

// GetSubway return sapporo subway operation info
func GetSubway() {
	values := url.Values{}
	values.Add("resource_id", "2daf777d-cb9b-4348-abc7-b69cd742cc12")
	values.Add("limit", "1")
	resp, err := http.Get("https://ckan.pf-sapporo.jp/api/action/datastore_search" + "?" + values.Encode())

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	ParseSubway(resp)
}

// ParseSubway http.Response to Parse Json
func ParseSubway(resp *http.Response) {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	data := new(model.Subway)

	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	fmt.Println(data)
}
