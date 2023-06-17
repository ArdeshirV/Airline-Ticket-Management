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

var (
	dbConn *gorm.DB

	db       string
	host     string
	port     string
	ssl      string
	timezone string
	user     string
	password string
)

func init() {
	db = config.Get(config.PostgresDB)
	host = config.Get(config.DatabaseHost)
	port = config.Get(config.DatabasePort)
	ssl = config.Get(config.PostgresSSL)
	timezone = config.Get(config.PostgresTimezone)
	user = config.Get(config.PostgresUser)
	password = config.Get(config.PostgresPassword)
}

func GetDSN() string {
	conStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
	// fmt.Printf("ConnectionString = \"%v\"\n", conStr)  // DEBUG: Present connection string
	return conStr
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
