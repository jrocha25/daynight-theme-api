package main

import (
	"log"
	"net/http"

	"daynight-theme.dev/api/lib"
	"daynight-theme.dev/api/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	godotenv.Load()

	// Checks if DB_URL environment variable is set
	_, err := lib.GetDBUrl()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := lib.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()
	log.Println("Database connection established!")

	log.Println("Seeding country info...")
	err = lib.SeedDatabase()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done seeding country info!")

	r := routes.MainRouter()

	log.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
