package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	service "../Service"
)

// Data for saving weather
type Data struct {
	Name    string
	Weather string
	Main    string
}

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Message from index"))
}

// Weather get weather
func Weather(w http.ResponseWriter, r *http.Request) {
	x := service.GetWeather()
	var result map[string]interface{}
	json.Unmarshal(x, &result)
	// main := fmt.Sprintf("%v", result["main"])
	// weather := fmt.Sprintf("%v", result["weather"])
	// name := fmt.Sprintf("%v", result["name"])
	// data := Data{Name: name, Weather: weather, Main: main}
	// xy, _ := json.Marshal(x)

	err := ioutil.WriteFile("./beans.txt", x, 0644)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(x)
}
