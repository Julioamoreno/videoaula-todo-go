package main

import (
	"fmt"
	"log"
	"net/http"
	"videoAulaGo/configs"
	"videoAulaGo/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Printf("Error to load configs")
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
