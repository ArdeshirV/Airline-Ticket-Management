package domain

type Ticket struct {
	ID        int
	Flight    *Flight
	Passenger *Passenger
	Payment   *Payment
	User      User
	StatusPay string
	Refund    bool
}
