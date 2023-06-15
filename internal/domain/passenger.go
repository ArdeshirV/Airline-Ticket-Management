package domain

import (
	"time"

	"gorm.io/gorm"
)

type Passenger struct {
	gorm.Model
	FirstName    string
	LastName     string
	NationalCode string
	Email        string
	Gender       Gender
	Phone        string
	BirthDate    time.Time
	Address      string
	UserID       uint
	Tickets      []Ticket
}
