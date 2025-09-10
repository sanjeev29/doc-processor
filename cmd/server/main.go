package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("server running at localhost:8000")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
