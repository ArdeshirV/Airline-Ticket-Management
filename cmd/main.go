package main

import (
	"fmt"
	"log"

	"github.com/the-go-dragons/final-project/internal/app"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
	_ "github.com/the-go-dragons/final-project/pkg/logger"
	"github.com/the-go-dragons/final-project/pkg/seeder"
)

func main() {
	database.CreateDBConnection()
	err := database.AutoMigrateDB()
	if err != nil {
		fmt.Print(err.Error())
	}
	app := app.NewApp()
	seeder.Run()
	log.Fatalln(app.Start(config.Get(config.PortMain)))
}
