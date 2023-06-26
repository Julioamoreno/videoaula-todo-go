package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"videoAulaGo/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error to get data: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
