package main

import (
	db "github.com/DevloperAmanSingh/news-api/internal/database"

	"log"
	"os"

	"github.com/DevloperAmanSingh/news-api/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	// Start the server
	app := router.SetupRouter()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DATABASE_NAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	db.ConnectDatabase(dbURL, dbName)
	defer db.DisconnectDatabase()
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
