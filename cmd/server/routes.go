package main

import (
	"net/http"

	"github.com/g3offrey/honeypot/pkg/catch"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
)

func registerRoutes(repos repositories, logger *httplog.Logger, r chi.Router) {
	var HONEYPOT_ROUTES = [...]string{
		"/wp-admin",
		"/administrator",
		"/admin",
	}

	for _, route := range HONEYPOT_ROUTES {
		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			err := catch.RegisterCatch(repos.catch, r)
			if err != nil {
				logger.Error(err.Error())
			}

			render.Status(r, http.StatusOK)
			render.PlainText(w, r, "catched")
		})
	}
}
