package main

import (
	"log"
	"net/http"

	"github.com/Jeet-baldha/codelink-backend/api"
)

func main() {

	router := api.SetupRoutes()

	log.Printf("🚀 Server running on http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}
