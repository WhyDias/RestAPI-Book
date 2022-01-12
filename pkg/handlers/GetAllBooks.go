package handlers

import (
	"RestAPI-Basic/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var book []models.Book

	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
