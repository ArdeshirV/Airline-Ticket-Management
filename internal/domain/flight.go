package domain

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNo        string
	DepartureID     uint
	Departure       Airport
	DestinationID   uint
	Destination     Airport
	DepartureTime   time.Time
	ArrivalTime     time.Time
	AirlineID       uint
	Airline         Airline
	FlightClass     FlightClass
	Price           int
	Capacity        int
	CancelCondition string
}
