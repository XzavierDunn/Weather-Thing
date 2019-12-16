package model

import (
    "log"
    "io/ioutil"
)

func saveWeather(x) err error{
    err := ioutil.WriteFile("./beans.txt", x, 0644)
	if err != nil {
		log.Fatal(err)
        return 
    return true
	}
}
