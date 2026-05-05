package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jsam1904/Proyecto1-Tracker-Api/db"
	"github.com/jsam1904/Proyecto1-Tracker-Api/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables directly")
	}

	db.Connect()
	defer db.DB.Close()

	r := router.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Bundesliga API running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
