package sample

import "net/http"

// Hello just say "Hello"
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(("Hello")))
}
