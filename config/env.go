package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig  struct {
	MONGOURI string
	JWT_KEY string
}

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig.MONGOURI = os.Getenv("MONGOURI")
	AppConfig.JWT_KEY = os.Getenv("JWT_KEY")
}
