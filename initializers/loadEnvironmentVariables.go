package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var GMAIL_APP_PASSWORD string

func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	GMAIL_APP_PASSWORD = os.Getenv("GMAIL_APP_PASSWORD")

	log.Println("Successfully loaded environment variables")
}
