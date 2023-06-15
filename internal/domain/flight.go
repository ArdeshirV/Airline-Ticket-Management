package domain

import "time"

type Flight struct {
	ID              int
	FlightNo        string
	Departure       Airport
	Destination     Airport
	DepartureTime   time.Time
	ArrivalTime     time.Time
	Airlines        Airline
	FlightClass     FlightClass
	Price           float64
	Capacity        int
	CancelCondition string
}
