package main

import (
	"os"

	"github.com/the-go-dragons/final-project/pkg/config"
)

func main() {
	config.LoadEnvVariables()
	println(os.Getenv("DATABASE_HOST")) // Example of using environment variable
}
