package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running")

	r := router.HandleRoutes()

	log.Fatal(http.ListenAndServe(":5000", r))
}
