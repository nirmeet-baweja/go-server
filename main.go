package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Fprintf used to write a fornatted string to the specified writer
	fmt.Fprintf(w, "welcome to the Home Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal((http.ListenAndServe(":3000", nil)))
}

func main() {
	handleRequests()
}
