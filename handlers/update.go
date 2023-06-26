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

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error to id parse: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error to json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		log.Printf("Error to update: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var resp map[string]any

	if rows != 1 {
		log.Printf("Error: Updated %d data", rows)

		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Data not been updated"),
		}

	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Data has been updated! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
