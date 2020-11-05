package main

import (
	"log"
	"net/http"
)

func main() {
	initLOGO()

	mux := initRouter()
	err := http.ListenAndServe(":8010", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func initLOGO() {
	println(`
             _            
 _ __   ___ | |_ ___  ___ 
| '_ \ / _ \| __/ _ \/ __|
| | | | (_) | ||  __/\__ \
|_| |_|\___/ \__\___||___/
`)
}
