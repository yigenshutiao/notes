package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8000"

func main() {

	initLOGO()

	if err := initDB(); err != nil {
		fmt.Println("init MySQL failed")
	}

	mux := initRouter()
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
