package main

import (
	"fmt"
	"net/http"
	"notes/logging"
)

const port = "8000"

func main() {

	initLOGO()

	logging.InitLog()

	fmt.Println("init DB")
	if err := initDB(); err != nil {
		logging.Fatal("init DB failed")
	}

	fmt.Println("init Cache")
	if err := initCache(); err != nil {
		logging.Fatal("init Cache failed")
	}

	fmt.Println("init HTTP Router")
	mux := initRouter()

	fmt.Println("start HTTP server")
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		logging.Fatal(err)
	}
}
