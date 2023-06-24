package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PayAmount     int64     `json:"payamount"`
	PayTime       time.Time `json:"paytime"`
	PaymentSerial string    `json:"paymentserial"`
	OrderID       uint      `json:"orderid"`
	Order         Order     `json:"order"`
}
