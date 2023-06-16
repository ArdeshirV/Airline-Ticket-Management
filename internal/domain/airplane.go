package domain

import "gorm.io/gorm"

type Airplane struct {
	gorm.Model
	Name      string
	AirlineID uint
	Airline   Airline
	Capacity  uint
}
