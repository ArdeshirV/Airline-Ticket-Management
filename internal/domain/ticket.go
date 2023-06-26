package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	FlightID      uint      `json:"-" gorm:"foreignKey:FlightID"`
	Flight        Flight    `json:"flight" gorm:"references:ID"`
	PassengerID   uint      `json:"-" gorm:"foreignKey:PassengerID"`
	Passenger     Passenger `json:"passenger" gorm:"references:ID"`
	PaymentID     uint      `json:"paymentid" gorm:"foreignKey:PaymentID"`
	Payment       Payment   `json:"payment" gorm:"references:ID"`
	UserID        uint      `json:"userid"  gorm:"foreignKey:UserID"`
	User          User      `json:"user" gorm:"references:ID"`
	PaymentStatus string    `json:"paymentstatus"`
	Refund        bool      `json:"refund"`
}
