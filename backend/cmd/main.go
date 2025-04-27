package main

import (
	"backend/handlers"
	"backend/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	logger.InitLogger()
	r := mux.NewRouter()
	r.HandleFunc("/items", handlers.Items).Methods("GET", "OPTIONS")
	r.HandleFunc("/items/{id}", handlers.ItemByID).Methods("GET", "OPTIONS")
	r.HandleFunc("/items/create", handlers.CreateItem).Methods("POST", "OPTIONS")
	r.HandleFunc("/items/update/{id}", handlers.UpdateItem).Methods("PUT", "OPTIONS")
	r.HandleFunc("/items/delete/{id}", handlers.DeleteItem).Methods("DELETE", "OPTIONS")

	port := ":50001"
	logger.InfOf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
