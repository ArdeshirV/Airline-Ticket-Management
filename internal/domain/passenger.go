package domain

import (
	"time"

	"gorm.io/gorm"
)

type Passenger struct {
	gorm.Model
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	NationalCode string    `json:"nationalcode"`
	Email        string    `json:"email"`
	Gender       Gender    `json:"gender"`
	Phone        string    `json:"phone"`
	BirthDate    time.Time `json:"birthdate"`
	Address      string    `json:"address"`
	UserID       uint      `json:"userid"`
	Tickets      []Ticket  `json:"tickets"`
}
