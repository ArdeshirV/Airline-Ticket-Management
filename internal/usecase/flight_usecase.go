package usecase

import (
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type FlightUseCase struct {
	flightRepo *persistence.FlightRepository
}

func NewFlightUseCase(flightRepo *persistence.FlightRepository) *FlightUseCase {
	return &FlightUseCase{
		flightRepo: flightRepo,
	}
}

func (fu FlightUseCase) IncreaseFlightCapacity(flight *domain.Flight) error {
	return fu.flightRepo.IncreaseFlightCapacity(flight)
}
