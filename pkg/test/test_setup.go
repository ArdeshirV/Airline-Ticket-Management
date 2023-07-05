package test

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
)

var rootPath = "."

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	rootPath = dir
}

func Setup() {
	config.Load(rootPath)
	database.Load()
	database.CreateDBConnection()
}
func SetupWithData() {
	Setup()
	loadTestData()
}
func Teardown() {
	clearTestData()
}

func loadTestData() {
	dat, err := os.ReadFile(rootPath + "/pkg/test/data/insert_data.sql")
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(string(dat))
}

func clearTestData() {
	dat, err := os.ReadFile(rootPath + "/pkg/test/data/remove_data.sql")
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(string(dat))
}
