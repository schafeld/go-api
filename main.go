package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the simple GO API!")
    fmt.Println("Endpoint Hit: root directory")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    fmt.Println("Started web server at http://localhost:8080")
    handleRequests()
}
