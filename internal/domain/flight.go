package domain

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNo          string
	DepartureID       uint
	Departure         Airport
	DestinationID     uint
	Destination       Airport
	DepartureTime     time.Time
	ArrivalTime       time.Time
	AirplaneID        uint
	Airplane          Airplane
	FlightClass       FlightClass
	Price             int
	RemainingCapacity int
	CancelCondition   string
	// AirlineID       uint
	// Airline         Airline  --> Already exists in Airplane
}
