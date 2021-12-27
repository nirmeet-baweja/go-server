package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content`
}

// global Articles array/slice to store dummy data
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	// Fprintf used to write a fornatted string to the specified writer
	fmt.Fprintf(w, "welcome to the Home Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// add route to get all the articles
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal((http.ListenAndServe(":3000", nil)))
}

func main() {
	Articles = []Article{
		{Title: "Article 1",
			Description: "Theme of the article",
			Content:     "The content of the article goes here."},
		{Title: "Article 1",
			Description: "Theme of the article",
			Content:     "The content of the article goes here."},
	}
	handleRequests()
}
