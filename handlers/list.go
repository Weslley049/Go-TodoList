package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Weslley049/todo-list/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Houve um erro ao obter os registros %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
