package routes

import (
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
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
