package main

import (
	"log"
	"net/http"
)

const port = "8000"

func main() {

	initLOGO()

	mux := initRouter()
	err := http.ListenAndServe(":"+port, mux)
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
