package main

import (
	"fmt"
	"net/http"

	"github.com/g3offrey/honeypot/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog/v2"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	repos, err := getRepositories(&cfg)
	if err != nil {
		panic(err)
	}

	logger := getLogger(&cfg)

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))

	registerRoutes(repos, logger, r)

	listenAddr := ":" + cfg.Port
	logger.Info(fmt.Sprintf("Listening on %s", listenAddr))
	err = http.ListenAndServe(listenAddr, r)
	if err != nil {
		panic(err)
	}
}
