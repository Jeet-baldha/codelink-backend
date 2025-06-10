package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jeet-baldha/codelink-backend/api"
	"github.com/Jeet-baldha/codelink-backend/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables first
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found or failed to load (fallback to OS env)")
	}

	// Connect to Mongo
	config.InitDB()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := api.SetupRoutes()
	log.Printf("🚀 Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
