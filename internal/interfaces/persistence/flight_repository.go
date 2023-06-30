package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type FlightRepository struct {
}

func NewFlightRepository() *FlightRepository {
	return &FlightRepository{}
}

func (a *FlightRepository) Create(input *domain.Flight) (*domain.Flight, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (a *FlightRepository) Update(input *domain.Flight) (*domain.Flight, error) {
	db, _ := database.GetDatabaseConnection()
	_, err := a.Get(int(input.ID))
	if err != nil {
		return nil, errors.New("the model doesnt exists")
	}
	tx := db.Where("id = ?", input.ID).Save(input)
	if tx.Error != nil {
		return input, tx.Error
	}
	tx.Commit()
	return input, nil
}

func (a *FlightRepository) Get(id int) (*domain.Flight, error) {
	var flight domain.Flight
	db, _ := database.GetDatabaseConnection()
	tx := db.First(&flight, id)
	if tx.Error != nil {
		return &flight, tx.Error
	}
	return &flight, nil
}

func (a *FlightRepository) GetAll() (*[]domain.Flight, error) {
	var flights []domain.Flight
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&flights)

	tx := db.Debug().Find(&flights)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &flights, nil
}

func (a *FlightRepository) Delete(id int) error {
	flight, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&flight)

	deleted := db.Debug().Delete(flight).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (a *FlightRepository) IncreaseFlightCapacity(flight *domain.Flight) error {
	flight.RemainingCapacity = flight.RemainingCapacity + 1
	_, err := a.Update(flight)

	if err != nil {
		return err
	}

	return nil
}
