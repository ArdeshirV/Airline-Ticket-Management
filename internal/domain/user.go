package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primarykey"`
	Username  string
	Password  string
	Email     string
	Phone     string
	CreatedAt time.Time
	// Role       Role
	// Passengers []*Passenger
}
