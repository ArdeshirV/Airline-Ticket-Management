package database

import (
	"fmt"
	"log"
	"time"

	model "github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var user string
var password string
var db string
var host string
var port string
var ssl string
var timezone string
var dbConn *gorm.DB

// To initialize db config
func init() { // todo: remove this init, use a loaded config object, or simply os.GetEnv in your GetDSN function
	user = config.GetEnv("POSTGRES_USER", "admin")
	password = config.GetEnv("POSTGRES_PASSWORD", "admin")
	db = config.GetEnv("POSTGRES_DB", "gormDb2")
	host = config.GetEnv("DATABASE_HOST", "127.0.0.1")
	port = config.GetEnv("DATABASE_PORT", "5432")
	ssl = config.GetEnv("POSTGRES_SSL", "disable")
	timezone = config.GetEnv("POSTGRES_TIMEZONE", "Asia/Tehran")
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
		&model.Airplane{},
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
