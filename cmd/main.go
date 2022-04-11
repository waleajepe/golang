package main

import (
	"crud/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	log.Println("Hey! API is running")
	http.ListenAndServe(":4000", router)

}
