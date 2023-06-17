package domain

import "gorm.io/gorm"

type Airplane struct {
	gorm.Model
	Name      string  `json:"name"`
	AirlineID uint    `json:"airlineid"`
	Airline   Airline `json:"airline"`
	Capacity  uint    `json:"capacity"`
}
