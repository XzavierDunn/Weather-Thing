package controller

import (
	"encoding/json"
	"net/http"

	service "../Service"
)

// Data for api
type Data struct {
	Coord      Coord
	weather    Weath
	base       string
	main       Main
	visibility int
	wind       Wind
	rain       Rain
	clouds     Cloud
	dt         int
	sys        Sys
	timezone   int
	id         int
	name       string
	cod        int
}

// Coord for coord
type Coord struct {
	lon int
	lat int
}

// Weath for weather
type Weath struct {
	id          int
	main        string
	description string
	icon        string
}

// Main for main
type Main struct {
	temp      int
	feelsLike int
	tempMin   int
	tempMax   int
	pressure  int
	humidity  int
}

// Wind for wind
type Wind struct {
	speed int
	deg   int
}

// Rain for rain
type Rain struct {
	oneh int
}

// Cloud for cloud
type Cloud struct {
	all int
}

// Sys for sys
type Sys struct {
	sysType int
	id      int
	country string
	sunrise int
	sunset  int
}

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Message from index"))
}

// Weather get weather
func Weather(w http.ResponseWriter, r *http.Request) {
	x := service.GetWeather()
	// fmt.Printf("%s", x)
	var data Data
	json.Unmarshal(x, &data)
	w.Write(x)
}
