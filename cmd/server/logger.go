package main

import (
	"github.com/g3offrey/honeypot/internal/config"
	"github.com/go-chi/httplog/v2"
)

func getLogger(cfg *config.Config) *httplog.Logger {
	return httplog.NewLogger("honeypot", httplog.Options{
		LogLevel:         cfg.SlogLevel(),
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
	})
}
