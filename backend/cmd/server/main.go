package main

import (
	"log"
	"net/http"

	"github.com/DJT24/foot/backend/internal/api"
)

func main() {
	mux := http.NewServeMux()

	// Create API server
	apiServer := api.NewServer()

	// Mount API server at /v1
	mux.Handle("/v1/", http.StripPrefix("/v1", apiServer))

	// Health check at root
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
