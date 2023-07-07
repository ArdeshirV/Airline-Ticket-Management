package main

import (
	"fmt"
	"log"

	"github.com/the-go-dragons/final-project/internal/app"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
	_ "github.com/the-go-dragons/final-project/pkg/logger"
	"github.com/the-go-dragons/final-project/pkg/seeder"
	"github.com/the-go-dragons/final-project/pkg/test"
)

func main() {
	config.Load(".")
	database.Load()
	database.CreateDBConnection()
	err := database.AutoMigrateDB()
	if err != nil {
		fmt.Print(err.Error())
	}
	test.SetupWithData()  // Load fake data
	defer test.Teardown() // Clean fake data
	app := app.NewApp()
	seeder.Run()
	log.Fatalln(app.Start(config.Config.Server.Port))
}
