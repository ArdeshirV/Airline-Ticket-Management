package main

import (
	"fmt"
	"log"
	"os"

	"github.com/the-go-dragons/final-project/internal/app"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
	_ "github.com/the-go-dragons/final-project/pkg/logger"
)

func main() {
	config.LoadEnvVariables()
	database.InitDB()
	err := database.DBConn.AutoMigrate(&domain.User{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	app := app.NewApp()
	log.Fatalln(app.Start(os.Getenv("PORT_MAIN")))
}
