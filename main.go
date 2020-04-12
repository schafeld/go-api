// main.go
package main

import (
	"encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Struct for articles
type Article struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Global array to mock article data
var Articles []Article

const serverPort string = "8080"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the simple GO API!")
	fmt.Println("Endpoint Hit: root directory")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "First Title", Desc: "Description of 1st article", Content: "Article One Content"},
		Article{Title: "Second Title", Desc: "Descriptionof 2nd article", Content: "Article Two Content"},
	}

	fmt.Println("Starting web server at http://localhost:8080")
	handleRequests()
}
