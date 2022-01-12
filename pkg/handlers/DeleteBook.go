package handlers

import (
	"RestAPI-Basic/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//Find the book by id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Delete that book
	h.DB.Delete(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode("Deleted")
}
