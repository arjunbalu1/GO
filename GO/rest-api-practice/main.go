package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	}
	fmt.Println("Endpoint hit: All articles")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
