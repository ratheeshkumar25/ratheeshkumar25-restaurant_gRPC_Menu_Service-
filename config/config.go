package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error to load env file")
	}
}
