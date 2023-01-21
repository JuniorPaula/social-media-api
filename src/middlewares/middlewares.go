package middlewares

import (
	"api/src/auth"
	"api/src/utils"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.TokenValidate(r); err != nil {
			utils.ResponseError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
