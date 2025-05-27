package main

import (
	"backend/handlers"
	"backend/logger"
	"log"
	"net/http"

	_ "backend/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Items API
// @version 1.0
// @description API for managing items
// @host localhost:50001
// @BasePath /
func main() {
	logger.InitLogger()
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/image", handlers.ServeImage).Methods("GET", "OPTIONS")
	r.HandleFunc("/items", handlers.Items).Methods("GET", "OPTIONS")
	r.HandleFunc("/items/{id}", handlers.ItemByID).Methods("GET", "OPTIONS")
	r.HandleFunc("/items/create", handlers.CreateItem).Methods("POST", "OPTIONS")
	r.HandleFunc("/items/update/{id}", handlers.UpdateItem).Methods("PUT", "OPTIONS")
	r.HandleFunc("/items/delete/{id}", handlers.DeleteItem).Methods("DELETE", "OPTIONS")

	port := ":50001"
	logger.InfOf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
