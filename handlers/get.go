package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Weslley049/todo-list/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, http.StatusText((http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		http.Error(w, http.StatusText((http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
