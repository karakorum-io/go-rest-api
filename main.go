package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : '/'")
	fmt.Fprintf(w, "Karakorum-io Golang REST API default endpoint version 1.0.0")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : '/all-articles'")
	articles := Articles{
		Article{Title: "Article1", Desc: "desc 1", Content: "Content 1"},
		Article{Title: "Article2", Desc: "desc 2", Content: "Content 2"},
	}

	json.NewEncoder(w).Encode(articles)
}

func allArticlesPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : '/all-articles'")
	articles := Articles{
		Article{Title: "Article1", Desc: "desc 1", Content: "Content 1"},
		Article{Title: "Article2", Desc: "desc 2", Content: "Content 2"},
	}

	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/all-articles", allArticles).Methods("GET")
	router.HandleFunc("/all-articles", allArticlesPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
	fmt.Println("Server running on PORT 8081")
}

func main() {
	handleRequests()
}
