package domain

import (
	"time"

	"gorm.io/gorm"
)

type FlightClass string

const (
	FirstClass    FlightClass = "First Class"
	BusinessClass FlightClass = "Business Class"
	EconomyClass  FlightClass = "Economic Class"
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
	Price             int64       `json:"price"`
	RemainingCapacity int         `json:"remainingcapacity"`
	CancelCondition   string      `json:"cancelcondition"`
}
