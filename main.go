package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitValiables()
	fmt.Println("Server running at port:", config.Port)

	r := router.HandleRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
