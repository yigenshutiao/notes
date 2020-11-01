package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	_, err := wr.Write([]byte("This is a awesome note app!"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hehe!")
}
