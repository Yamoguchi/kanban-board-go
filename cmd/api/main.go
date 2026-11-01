package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		fmt.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler)

	log.Println("Listening on port 8080")

	if err := http.ListenAndServe("127.0.0.1:8080", mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	} else {
		fmt.Println("Server was stopped")
	}
}
