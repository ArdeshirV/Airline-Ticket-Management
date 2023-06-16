package domain

import (
	"time"
)

type User struct {
	ID         int
	Username   string
	Password   string
	Email      string
	Phone      string
	CreatedAt  time.Time
	RoleID     uint
	Role       Role
	Passengers []Passenger
	IsLoginRequired bool `gorm:"default:false"`
}
