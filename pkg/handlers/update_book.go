package handlers

import (
	"crud/pkg/mocks"
	"crud/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading body")
	}
	defer r.Body.Close()
	// unmarshal body
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)
	// iterate over mock books
	for index, book := range mocks.Books {
		if book.Id == id {
			book.Author = updatedBook.Author
			book.Title = updatedBook.Title
			book.Desc = updatedBook.Desc
			// update book
			mocks.Books[index] = book
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("updated")
			break

		}
	}
	

}
