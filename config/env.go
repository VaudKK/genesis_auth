package config

import (
	"os"
)

var AppConfig struct {
	MONGOURI string
	JWT_KEY  string
}

func LoadConfig() {
	AppConfig.MONGOURI = os.Getenv("MONGOURI")
	AppConfig.JWT_KEY = os.Getenv("JWT_KEY")
}
