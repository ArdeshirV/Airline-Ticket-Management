package domain

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNo          string      `json:"flightno"`
	DepartureID       uint        `json:"departureid"`
	Departure         Airport     `json:"departure"`
	DestinationID     uint        `json:"destinationid"`
	Destination       Airport     `json:"destination"`
	DepartureTime     time.Time   `json:"departuretime"`
	ArrivalTime       time.Time   `json:"arrivaltime"`
	AirplaneID        uint        `json:"airplaneid"`
	Airplane          Airplane    `json:"airplane"`
	FlightClass       FlightClass `json:"flightclass"`
	Price             int         `json:"price"`
	RemainingCapacity int         `json:"remainingcapacity"`
	CancelCondition   string      `json:"cancelcondition"`
	// AirlineID       uint
	// Airline         Airline  --> Already exists in Airplane
}
