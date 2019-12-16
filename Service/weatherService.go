package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url = os.Getenv("API_URL")

//GetWeather send request to api
func GetWeather() []byte {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	json.Marshal(body)
	return body
}
