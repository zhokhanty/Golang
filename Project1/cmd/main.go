package main

import (
	"fmt"
	"go-project/api"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.HomeHandler).Methods("GET")
	r.HandleFunc("/captain", api.GetCaptain).Methods("GET")
	r.HandleFunc("/players", api.GetPlayersList).Methods("GET")
	r.HandleFunc("/players/{name}", api.GetPlayersDetails).Methods("GET")
	r.HandleFunc("/health", api.HealthCheck).Methods("GET")

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
