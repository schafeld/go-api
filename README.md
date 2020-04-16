# go-api

A simple API with Go.

## Based on Inspiration from

https://tutorialedge.net/golang/creating-restful-api-with-golang/

## Additional Go components/modules

[Gorilla/Mux Router](https://github.com/gorilla/mux)

Installation:

    go get -u github.com/gorilla/mux

## Testing API locally

### Read All

General response reading all articles as JSON:

    curl http://localhost:8080/articles

### Read One

Getting specific article as JSON by id:

    curl http://localhost:8080/article/2

### Write One

Posting (writing) a sample new article:

    curl --header "Content-Type: application/json" -d "{\"Id\":\"3\",\"Title\":\"Newly Created Post\",\"desc\":\"The description for my new post\",\"content\":\"my articles content\" }" http://localhost:8080/article

### Delete One

curl -X DELETE http://localhost:8080/article/3
