package domain

import "time"

type Payment struct {
	ID            int
	payAmount     int
	payTime       time.Time
	paymentSerial string
}
