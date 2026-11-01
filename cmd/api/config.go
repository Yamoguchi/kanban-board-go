package main

import (
	"errors"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	HTTPAddress       string        `envconfig:"ADDRESS" default:"127.0.0.1:8080"`
	ReadHeaderTimeout time.Duration `envconfig:"READ_HEADER_TIMEOUT" default:"5s"`
	ReadTimeout       time.Duration `envconfig:"READ_TIMEOUT" default:"10s"`
	WriteTimeout      time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
	IdleTimeout       time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s"`
}

func (cfg *ServerConfig) LoadFromEnv() error {
	err := envconfig.Process("HTTP", cfg)
	if err != nil {
		return err
	}

	if cfg.HTTPAddress == "" {
		return errors.New("HTTP_ADDRESS must not be empty")
	}
	if cfg.ReadHeaderTimeout <= 0 {
		return errors.New("HTTP_READ_HEADER_TIMEOUT must be positive")
	}
	if cfg.ReadTimeout <= 0 {
		return errors.New("HTTP_READ_TIMEOUT must be positive")
	}
	if cfg.WriteTimeout <= 0 {
		return errors.New("HTTP_WRITE_TIMEOUT must be positive")
	}
	if cfg.IdleTimeout <= 0 {
		return errors.New("HTTP_IDLE_TIMEOUT must be positive")
	}
	return nil
}
