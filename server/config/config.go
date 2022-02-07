package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Init - Initializing the Configuration via .env import
func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No env file found.")
	}

	port, portExists := os.LookupEnv("PORT")
	if portExists {
		fmt.Println("Localhost mount port (point postman here): " + port)
	} else {
		log.Fatal(".env vars not defined. \n`PORT` is required. \nSee .env.template for reference")
	}

	swApiEndpoint, swApiEndpointExists := os.LookupEnv("SW_API_ENDPOINT")
	if swApiEndpointExists {
		fmt.Println("Using SWAPI endpoint " + swApiEndpoint)
	} else {
		log.Fatal("SW_API_ENDPOINT required")
	}
}
