package config

import (
	"log"
	"os"

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

// This function can be used to get ENV Var with default value
func GetEnv(key, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
