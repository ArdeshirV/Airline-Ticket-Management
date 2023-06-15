package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	FlightID      uint
	Flight        Flight
	PassengerID   uint
	Passenger     Passenger
	PaymentID     uint
	Payment       Payment
	UserID        uint
	User          User
	PaymentStatus string
	Refund        bool
}
