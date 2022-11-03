package rotes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rote represents all api rotes
type Rote struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// Configure set all rotes inside of the router
func Configure(r *mux.Router) *mux.Router {
	rotes := rotesUsers
	rotes = append(rotes, rtlogin)

	for _, rote := range rotes {

		if rote.RequireAuthentication {

			r.HandleFunc(rote.URI, middlewares.Logger(middlewares.Authenticate(rote.Function))).Methods(rote.Method)

		} else {

			r.HandleFunc(rote.URI, middlewares.Logger(rote.Function)).Methods(rote.Method)
		}
	}
	return r
}
