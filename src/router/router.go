package router

import "github.com/gorilla/mux"

func HandleRoutes() *mux.Router {
	return mux.NewRouter()
}
