package usecase

import (
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type BookingAction struct {
	flightRepo    *persistence.FlightRepository
	passengerRepo *persistence.PassengerRepository
	ticketRepo    *persistence.TicketRepository
}

func NewBookingAction(flightRepo *persistence.FlightRepository,
	ticketRepo *persistence.TicketRepository) *Booking {
	return &Booking{
		flightRepo: flightRepo,
		ticketRepo: ticketRepo,
	}
}

func (b Booking) GetReservedTickets(userId uint) (*[]domain.Ticket, error) {
	return b.ticketRepo.GetByUserId(userId)
}

func (b Booking) GetAllReservedTickets() (*[]domain.Ticket, error) {
	return b.ticketRepo.GetAllNotArrivedByUserId()
}

func (b Booking) GetCancelledTickets(userId uint) (*[]domain.Ticket, error) {
	return b.ticketRepo.GetCancelledByUserId(userId)
}

func (b Booking) CancelTicket(ticketId int) (*domain.Ticket, error) {
	return b.ticketRepo.CancelTicket(ticketId)
}

// TODO: Move the logic of '/flights' end-point here
