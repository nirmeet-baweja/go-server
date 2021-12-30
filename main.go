package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// update the Article struct to include Id
type Article struct {
	Id          string `json:"id"`
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

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// find the article matching the Id
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// handling requests using the third party router - gorilla/mux
func handleRequests() {
	// create  a new instance of mux router
	router := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with router.HandleFunc
	router.HandleFunc("/", homePage)
	// add route to get all the articles
	router.HandleFunc("/articles", returnAllArticles)
	// add route to get a specific article based on article's ID
	router.HandleFunc("/articles/{id}", returnSingleArticle)
	// pass the router as a second argument to the http.ListenAndServe function
	log.Fatal((http.ListenAndServe(":3000", router)))
}

func main() {
	Articles = []Article{
		{Id: "1",
			Title:       "Article 1",
			Description: "Theme of the article",
			Content:     "The content of the article goes here."},
		{Id: "2",
			Title:       "Article 1",
			Description: "Theme of the article",
			Content:     "The content of the article goes here."},
	}
	handleRequests()
}
