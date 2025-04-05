package main

import (
	"backend/handlers"
	"backend/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	logger.InitLogger()
	r := mux.NewRouter()
	r.HandleFunc("/items", handlers.Items).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.ItemByID).Methods("GET")
	r.HandleFunc("/items/create", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/items/update/{id}", handlers.UpdateItem).Methods("POST")
	r.HandleFunc("/items/delete/{id}", handlers.DeleteItem).Methods("DELETE")

	port := ":50001"
	logger.InfOf("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
