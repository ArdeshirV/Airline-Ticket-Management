package main

import (
	"os"
	"log"
	_ "github.com/the-go-dragons/final-project/pkg/logger"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/internal/app"
)

func main() {
	config.LoadEnvVariables()
	app := app.NewApp()
	log.Fatalln(app.Start(os.Getenv("PORT_MAIN")))
}
