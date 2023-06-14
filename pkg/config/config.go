package config

import (
	"log"
	"github.com/joho/godotenv"
)

func init() {
	// TODO: Should create a thread-safe singleton here to handle configurations
}

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
