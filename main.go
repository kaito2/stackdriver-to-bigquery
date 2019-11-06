package sample

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/logging"
)

const project = "kaito2"

// User user info
type User struct {
	// Name user name
	Name string `json:"name"`
}

// Hello just say "Hello"
func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Get logging client
	client, err := logging.NewClient(ctx, project)
	if err != nil {
		log.Fatal("Failed to get logging client")
	}

	// Parse query string
	names, ok := r.URL.Query()["name"]
	var name string
	if !ok || len(names[0]) < 1 {
		name = "World"
	} else {
		name = names[0]
	}
	user := User{Name: name}

	// Get logger
	logger := client.Logger("sample-logger-1")

	// logging
	logger.Log(logging.Entry{
		Payload:  user,
		Severity: logging.Info,
		Labels:   map[string]string{"bigquery-export-type": "sample"},
	})

	logger.Log(logging.Entry{
		Payload:  "warning",
		Severity: logging.Warning,
	})

	logger.Log(logging.Entry{
		Payload:  "error",
		Severity: logging.Error,
	})

	// Validate http method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET method is allowed"))
	}

	msg := fmt.Sprintf("Hello %s", name)
	w.Write([]byte(msg))
}
