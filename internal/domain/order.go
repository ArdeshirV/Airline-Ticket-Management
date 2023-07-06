package domain

import "gorm.io/gorm"

type OrderStatus int

const (
	PENDING OrderStatus = iota
	PAID
	DELIVERED
)

type Order struct {
	gorm.Model
	OrderNum   string
	Amount     int64
	FlightID   uint
	Flight     Flight
	Status     OrderStatus
	OrderItems []OrderItem
	UserID     uint
}

type OrderItem struct {
	gorm.Model
	OrderID     uint
	Order       Order
	PassengerID uint
	Passenger   Passenger
}
