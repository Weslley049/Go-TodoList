package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Weslley049/todo-list/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, http.StatusText((http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		http.Error(w, http.StatusText((http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos mais de um registros : %d", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "registro removido com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
