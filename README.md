# go-api

A simple API with Go.

## Based on Inspiration from

https://tutorialedge.net/golang/creating-restful-api-with-golang/
and
https://www.youtube.com/watch?v=SonwZ6MF5BE&feature=youtu.be 
https://github.com/bradtraversy/go_restapi 

## Additional Go components/modules

[Gorilla/Mux Router](https://github.com/gorilla/mux)

Installation:

    go get -u github.com/gorilla/mux

## Testing API locally

### Read All (GET)

General response reading all articles as JSON:

    curl http://localhost:8080/articles

### Read One (GET)

Getting specific article as JSON by id:

    curl http://localhost:8080/article/2

### Write One (POST)

Posting (writing) a sample new article:

    curl --header "Content-Type: application/json" -d "{\"Id\":\"3\",\"Title\":\"Newly Created Post\",\"desc\":\"The description for my new post\",\"content\":\"my articles content\" }" http://localhost:8080/article

### Delete One (DELETE)

curl -X DELETE http://localhost:8080/article/3

### Update One (PUT)

Updating (i.e. overwriting) an existing article (Important: Append id parameter in URL):

    curl -X PUT -H "Content-Type: application/json" -d "{\"Id\":\"3\",\"Title\":\"Updated Post\",\"desc\":\"This is the updated description\",\"content\":\"my updated articles content\" }" http://localhost:8080/article/3
