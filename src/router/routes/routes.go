package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI               string
	Method            string
	Function          func(http.ResponseWriter, *http.Request)
	hasAuthentication bool
}

func SetupRoutes(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoutes)

	for _, route := range routes {

		if route.hasAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentication(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
