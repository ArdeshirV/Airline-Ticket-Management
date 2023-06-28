package usecase

import (
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type TicketUsecase struct {
	ticketRepository *persistence.TicketRepository
}

func NewTicketUsecase(repository *persistence.TicketRepository) *TicketUsecase {
	return &TicketUsecase{
		ticketRepository: repository,
	}
}

func (tu *TicketUsecase) CreateTicket(ticket *domain.Ticket) (*domain.Ticket, error) {
	return tu.ticketRepository.Create(ticket)
}
