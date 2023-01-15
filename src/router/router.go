package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func HandleRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.SetupRoutes(r)
}
