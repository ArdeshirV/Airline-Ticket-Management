package domain

import _ "gorm.io/gorm"

type Airplane struct {
	ID        uint
	Name      string
	AirlineID uint
	Airline   Airline
	Capacity  uint
}
