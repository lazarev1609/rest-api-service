package main

import (
	"log"
	"net/http"
	"rest-api-service/api"
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
