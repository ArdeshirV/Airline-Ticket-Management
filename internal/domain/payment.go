package domain

import "time"

type Payment struct {
	ID            int
	PayAmount     int
	PayTime       time.Time
	PaymentSerial string
}
