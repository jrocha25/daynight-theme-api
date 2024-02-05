package main

import (
	"log"
	"net/http"

	"daynight-theme.dev/api/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	godotenv.Load()

	r := routes.MainRouter()

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
