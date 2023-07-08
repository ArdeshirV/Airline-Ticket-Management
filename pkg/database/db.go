package database

import (
	"fmt"
	models "github.com/the-go-dragons/final-project/internal/domain"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB

	db       string
	host     string
	port     int
	ssl      string
	timezone string
	user     string
	password string
)

func Load() {
	db = config.Config.Database.Name
	host = config.Config.Database.Host
	port = config.Config.Database.Port
	ssl = config.Config.Database.Ssl
	timezone = config.Config.Database.Timezone
	user = config.Config.Database.User
	password = config.Config.Database.Password
}

func GetDSN() string {
	conStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, db, port, ssl, timezone)
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

	err = db.AutoMigrate(
		&models.Airplane{},
		&models.Airplane{},
		&models.Airport{},
		&models.City{},
		&models.Flight{},
		&models.Order{},
		&models.Passenger{},
		&models.Payment{},
		&models.Role{},
		&models.Ticket{},
		&models.User{},
	)
	if err != nil {
		return err
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
	conn, err := GetDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := pg.WithInstance(sqlDB, &pg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://./pkg/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	migrate.Up()
	return err
}
