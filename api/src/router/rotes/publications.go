package rotes

import (
	"api/src/controllers"
	"net/http"
)

var PublicationRotes = []Rote{

	{
		URI:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: false,
	},
	{
		URI:                   "/publications",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublications,
		RequireAuthentication: false,
	},
	{
		URI:                   "/publications/{publicationID}",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublicarion,
		RequireAuthentication: false,
	},
	{
		URI:                   "/publications/{publicationID}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequireAuthentication: false,
	},
	{
		URI:                   "/publications/{publicationID}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequireAuthentication: false,
	},
}
