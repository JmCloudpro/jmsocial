package rotes

import (
	"api/src/controllers"
	"net/http"
)

var rtlogin = Rote{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequireAuthentication: false,
}
