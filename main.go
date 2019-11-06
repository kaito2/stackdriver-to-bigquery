package sample

import (
	"fmt"
	"net/http"
)

// Hello just say "Hello"
func Hello(w http.ResponseWriter, r *http.Request) {
	// Validate http method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET method is allowed"))
	}
	
	// Parse query string
	names, ok := r.URL.Query()["name"]
	var name string
	if !ok || len(names[0]) < 1 {
		name = "World"
	} else {
		name = names[0]
	}

	msg := fmt.Sprintf("Hello %s", name)
	w.Write([]byte(msg))
}
