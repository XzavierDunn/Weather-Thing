package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url = os.Getenv("API_URL")

// Data struct
type Data struct {
	coord      []Coord
	weather    []Weather
	base       string
	main       []Main
	visibility int
}

// Coord struct
type Coord struct {
	lon int
	lat int
}

// Weather struct
type Weather struct {
	id          int
	main        string
	description string
	icon        string
}

// Main struct
type Main struct {
	temp       int
	feels_like int
	temp_min   int
	temp_max   int
	pressure   int
	humidity   int
}

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

	data := Data{}
	x := json.Unmarshal(body, &data)

	fmt.Printf("%v", x)

	// json.Marshal(body)
	return body
}
