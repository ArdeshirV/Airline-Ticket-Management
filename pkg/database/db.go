package database

import (
	"fmt"
	"os"

	"github.com/the-go-dragons/final-project/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() error {
	var err error

	// Get the postgres connection data
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"

	// Connect to postgres
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Connected to database")

	// Migrate the tables
	err = DBConn.AutoMigrate(&domain.User{})
	fmt.Println("Migrated the tables")

	return err
}
