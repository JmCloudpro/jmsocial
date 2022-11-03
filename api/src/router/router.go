package router

import (
	"api/src/router/rotes"

	"github.com/gorilla/mux"
)

// Generate  returns a router with routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return rotes.Configure(r)
}
