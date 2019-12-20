package main

import (
	"fmt"
	"log"
	"net/http"

	controller "./Controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Index)
	mux.HandleFunc("/weather", controller.Weather)

	fmt.Println("Listening... localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
