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

	containerAPIPort, containerAPIPortExists := os.LookupEnv("CONTAINER_API_PORT")
	hostAPIPort, hostAPIPortExists := os.LookupEnv("HOST_API_PORT")
	if containerAPIPortExists && hostAPIPortExists {
		os.Setenv("containerAPIPort", containerAPIPort)
		os.Setenv("hostAPIPort", hostAPIPort)
		fmt.Println("Container expose port: " + os.Getenv("containerAPIPort"))
		fmt.Println("Localhost mount port (point postman here): " + os.Getenv("hostAPIPort"))
	} else {
		log.Fatal(".env vars not defined. \nCONTAINER_API_PORT and HOST_API_PORT are both required. \nSee .env.template for reference")
	}

	swApiEndpoint, swApiEndpointExists := os.LookupEnv("SW_API_ENDPOINT")
	if swApiEndpointExists {
		fmt.Println("Using SWAPI endpoint " + swApiEndpoint)
	} else {
		log.Fatal("SW_API_ENDPOINT required")
	}
}
