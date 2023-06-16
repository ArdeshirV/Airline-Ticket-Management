package database

import (
	"fmt"
	"log"
	"os"
	"time"

	model "github.com/the-go-dragons/final-project/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// This function can be used to get ENV Var with default value
func getenv(key, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}

var user string
var password string
var db string
var host string
var port string
var ssl string
var timezone string
var dbConn *gorm.DB

// To initialize db config
func init() {
	user = getenv("POSTGRES_USER", "admin")
	password = getenv("POSTGRES_PASSWORD", "admin")
	db = getenv("POSTGRES_DB", "gormDb2")
	host = getenv("DATABASE_HOST", "127.0.0.1")
	port = getenv("DATABASE_PORT", "5432")
	ssl = getenv("POSTGRES_SSL", "disable")
	timezone = getenv("POSTGRES_TIMEZONE", "Asia/Tehran")
}

func GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
}

func CreateDBConnection() error {
	// Close the existing connection if open
	if dbConn != nil {
		CloseDBConnection(dbConn)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db
	return err
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

func CloseDBConnection(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
}

func AutoMigrateDB() error {
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		return connErr
	}
	// Add new models here
	err := db.AutoMigrate(
		&model.Airline{},
		&model.Airport{},
		&model.City{},
		&model.Flight{},
		&model.Passenger{},
		&model.Payment{},
		&model.Role{},
		&model.Ticket{},
		&model.User{},
	)
	return err
}
