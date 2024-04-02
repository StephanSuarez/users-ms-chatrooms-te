package http

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var projectIDPub string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectIDPub = os.Getenv("PROJECT_ID")
}
