// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct for articles
type Article struct {
	Id      string `json:"Id"`
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

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprintf(w, "Key: "+key)

	// Loop over all of Articles
	// if article.Id equals passed-in key
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	// get body of POST request
	// return string response containing request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update global Articles array with new article
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	// fmt.Fprintf(w, "%+v", string(reqBody))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	// get parameters from URL
	vars := mux.Vars(r)
	// extract `id` of article to delete
	id := vars["id"]

	// loop through all articles
	for index, article := range Articles {
		// delete article with matching id
		if article.Id == id {
			// update Articles array without deleted article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

// TODO: Make this work
func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	// w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.Id = params["id"]
			Articles = append(Articles, article)
			json.NewEncoder(w).Encode(article)
			return
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	// Beware this post-article route needs to come first
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "First Title", Desc: "Description of 1st article", Content: "Article One Content"},
		Article{Id: "2", Title: "Second Title", Desc: "Description of 2nd article", Content: "Article Two Content"},
	}

	fmt.Println("Starting web server at http://localhost:8080")
	handleRequests()
}
