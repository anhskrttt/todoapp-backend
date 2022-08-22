package main

import (
	"fmt"
	"go-server/router"
	"log"
	"net/http"
)

func main() {
	// Init Router.
	r := router.Router()

	// Start server.
	fmt.Println("Listening and serving.....")
	log.Fatal(http.ListenAndServe(":8080", r))
}
<<<<<<< HEAD
=======
