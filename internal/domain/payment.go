package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PayAmount     int
	PayTime       time.Time
	PaymentSerial string
}
