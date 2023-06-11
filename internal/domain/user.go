package domain

import "time"

type User struct {
	ID         int
	Username   string
	Password   string
	Email      string
	Phone      string
	CreatedAt  time.Time
	Role       Role
	Passengers []*Passenger
}
