package main

import (
	"net/http"
	"notes/logic"
)

func initRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(logic.Hello))

	return mux
}
