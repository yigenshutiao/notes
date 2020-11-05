package bind

import (
	"net/http"
)

const (
	Content = "Content-Type"
)

type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

func Bind(r *http.Request, target interface{}) error {
	contentType := getContentType(r)
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	return nil
}

func GetBinder(method, contentType string) Binding {
	if method == "GET" {
		return
	}
}

func getContentType(r *http.Request) string {
	if content, ok := r.Header[Content]; ok {
		cont := content[0]
		for i, char := range cont {
			if char == ' ' || char == ';' {
				return cont[:i]
			}
		}
		return cont
	}
	return ""
}
