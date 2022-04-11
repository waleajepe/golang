package handlers

import (
	"crud/pkg/mocks"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// read id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//iterate over books
	for index, book := range mocks.Books {
		// delete book index where book.id equals vars id
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
