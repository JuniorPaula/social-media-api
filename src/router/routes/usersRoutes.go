package routes

import "net/http"

var usersRoutes = []Route{
	{
		URI:               "/users",
		Method:            http.MethodPost,
		Function:          func(w http.ResponseWriter, r *http.Request) {},
		hasAuthentication: false,
	},
	{
		URI:               "/users",
		Method:            http.MethodGet,
		Function:          func(w http.ResponseWriter, r *http.Request) {},
		hasAuthentication: false,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodGet,
		Function:          func(w http.ResponseWriter, r *http.Request) {},
		hasAuthentication: false,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodPut,
		Function:          func(w http.ResponseWriter, r *http.Request) {},
		hasAuthentication: false,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodDelete,
		Function:          func(w http.ResponseWriter, r *http.Request) {},
		hasAuthentication: false,
	},
}
