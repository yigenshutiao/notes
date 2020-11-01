package logic

import "net/http"

func Hello(wr http.ResponseWriter, r *http.Request) {
	_, err := wr.Write([]byte("This is a awesome note app!"))
	if err != nil {
		return
	}
}
