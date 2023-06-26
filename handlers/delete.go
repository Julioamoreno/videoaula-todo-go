package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"videoAulaGo/models"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error to id parse: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Error to delete: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows != 1 {
		log.Printf("Error: Delete %d data", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Data has been delete! ID: %d", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
