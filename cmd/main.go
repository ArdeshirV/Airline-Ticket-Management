package main

import (
	"fmt"
	"log"
	"os"

	"github.com/the-go-dragons/final-project/internal/app"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
	_ "github.com/the-go-dragons/final-project/pkg/logger"
)

func main() {
	config.LoadEnvVariables()
	database.CreateDBConnection()
	err := database.AutoMigrateDB()
	if err != nil {
		fmt.Printf(err.Error())
	}
	app := app.NewApp()
	log.Fatalln(app.Start(os.Getenv("PORT_MAIN")))
}
