package model

import (
//    "os"
//    "os/exec"
	"io/ioutil"
	"log"
)

// SaveWeather to file
func SaveWeather(x []byte) {
	err := ioutil.WriteFile("./beans.txt", x, 0644)
	if err != nil {
		log.Fatal(err)
	}

//    file := os.Getenv("pyFile")
//    c := exec.Command(file)
//    if err := c.Run(); err != nil {
//        log.Fatal(err)
//    }
}
