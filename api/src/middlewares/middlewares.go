package middlewares

import (
	"api/src/answers"
	"api/src/authentication"
	"log"
	"net/http"
)

// Logger writes rquest informations on term
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate  receives parameters (w http.ResponseWriter, r http.Request) and returns the same HandlerFunc
// check is the user is auhenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := authentication.ValidateToken(r); err != nil {

			answers.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}

}
