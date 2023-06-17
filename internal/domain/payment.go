package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PayAmount     int       `json:"payamount"`
	PayTime       time.Time `json:"paytime"`
	PaymentSerial string    `json:"paymentserial"`
}
