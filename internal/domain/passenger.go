package domain

type Passenger struct {
	ID           int
	FirstName    string
	LastName     string
	NationalCode string
	Email        string
	Gender       Gender
	Phone        string
	Age          uint
	Address      string
	Tickets      []*Ticket
}
