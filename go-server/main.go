package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World")

	// Init Router.
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/", HomePageHandler).Methods("GET")

	// Start server.
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	welcome := "Welcome to my homepage"
	json.NewEncoder(w).Encode(&welcome)
}

// type ToDo struct {
// 	ID int `json: "id"`
// 	Name string `json: "name"`
// 	Description string `json: "description"`
// 	Status bool `json: "status"`
// }
