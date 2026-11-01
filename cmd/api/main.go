package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	HTTPAddress string `envconfig:"ADDRESS" default:"127.0.0.1:8080"`
}

func (cfg *ServerConfig) LoadFromEnv() error {
	err := envconfig.Process("HTTP", cfg)
	if err != nil {
		return err
	}
	if cfg.HTTPAddress == "" {
		return errors.New("HTTP_ADDRESS must not be empty")
	}
	return nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		fmt.Println(err)
	}
}

func main() {
	//Read configuration from environment
	var cfg ServerConfig
	if err := cfg.LoadFromEnv(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler)

	log.Println("Listening on", cfg.HTTPAddress)

	if err := http.ListenAndServe(cfg.HTTPAddress, mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	} else {
		fmt.Println("Server was stopped")
	}
}
