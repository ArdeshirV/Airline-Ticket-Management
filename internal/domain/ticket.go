package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	FlightID      uint      `json:"flightid"`
	Flight        Flight    `json:"flight"`
	PassengerID   uint      `json:"passengerid"`
	Passenger     Passenger `json:"passenger"`
	PaymentID     uint      `json:"paymentid"`
	Payment       Payment   `json:"payment"`
	UserID        uint      `json:"userid"`
	User          User      `json:"user"`
	PaymentStatus string    `json:"paymentstatus"`
	Refund        bool      `json:"refund"`
}
