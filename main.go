package main

import (
	"log"
	"net/http"
)

func main() {
	mux := initRouter()
	err := http.ListenAndServe(":8010", mux)
	if err != nil {
		log.Fatal(err)
	}
}
