package main

import (
	"log"
	"net/http"
	"github.com/DJT24/foot/backend/internal/api"
)

func main() {
	srv := api.NewServer()
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
