package common

import (
	"github.com/yigenshutiao/binding"
	"net/http"
)

func getContentType(r *http.Request) string {

	ct := ""
	if contentType, ok := r.Header["Content-Type"]; ok {
		ct = contentType[0]
	}

	for i, char := range ct {
		if char == ' ' || char == ';' {
			return ct[:i]
		}
	}

	return ct
}

func Bind(r *http.Request, dst interface{}) error {

	contentType := getContentType(r)

	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	bind := binding.Default(r.Method, contentType)

	return bind.Bind(r, dst)
}
