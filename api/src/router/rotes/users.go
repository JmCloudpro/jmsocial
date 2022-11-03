package rotes

import (
	"api/src/controllers"
	"net/http"
)

var rotesUsers = []Rote{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},

	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.SearchUsers,

		RequireAuthentication: false,
	},

	//get a user by ID
	{
		URI:      "/users/{userID}",
		Method:   http.MethodGet,
		Function: controllers.SearchUser,

		RequireAuthentication: true,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,

		RequireAuthentication: true,
	},

	{
		URI:                   "/users/{userID}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},

	{
		URI:                   "/users/{userID}/follow",
		Method:                http.MethodPost,
		Function:              controllers.Followuser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userID}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.Unollowuser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userID}/followers",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userID}/updatepassword",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePawword,
		RequireAuthentication: true,
	},
}
