package controller

import (
	"encoding/json"
	"net/http"

	model "../Model"
	service "../Service"
)

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Message from index"))
}

// Weather get weather
func Weather(w http.ResponseWriter, r *http.Request) {
	x := service.GetWeather()
	var result map[string]interface{}
	json.Unmarshal(x, &result)
	model.SaveWeather(x)
    w.Write(x)
}
