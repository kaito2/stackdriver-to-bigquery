package sample

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zapdriver.NewProduction()
	if err != nil {
		log.Fatalf("Failed to get logger: %v", err)
	}
}

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

	logger.Debug("debug")
	logger.Info(name)
	logger.Info(name, zapdriver.Label("bigquery-export", "true"))
	logger.Warn("warn")
	logger.Error("warn")
	log.Println("log.Println")

	msg := fmt.Sprintf("Hello %s", name)
	w.Write([]byte(msg))
}
