package handlers

import (
	"crud/pkg/mocks"
	"crud/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// read to request body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal("Error reading body")
	}
	// unamrshal the body to Book structure
	var book models.Book
	json.Unmarshal(body, &book)
	// Append to Book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)
	// send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Book created")

}
