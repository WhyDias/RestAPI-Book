package handlers

import (
	"RestAPI-Basic/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//Find book from id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Id id is are equal send book as response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}