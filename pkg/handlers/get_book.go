package handlers

import (
	"crud/pkg/mocks"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Read id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// Iterterate over all mock books
	for _, book := range mocks.Books {
	// id ids are equal, send a response
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
