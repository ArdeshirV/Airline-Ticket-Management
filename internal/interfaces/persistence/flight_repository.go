package persistence

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type FlightRepository interface {
	Create(input *domain.Flight) (*domain.Flight, error)
	Update(input *domain.Flight) (*domain.Flight, error)
	Get(id int) (*domain.Flight, error)
	GetAll() (*[]domain.Flight, error)
	Delete(id int) error
}
type FlightRepositoryImp struct {
}

func NewFlightRepository() FlightRepository {
	return &FlightRepositoryImp{}
}

func (a FlightRepositoryImp) Create(input *domain.Flight) (*domain.Flight, error) {
	var flight domain.Flight
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&flight)

	checkFlightExist := db.Debug().First(&flight, "ID = ?", input.ID)

	if checkFlightExist.RowsAffected > 0 {
		return &flight, errors.New(strconv.Itoa(http.StatusConflict))
	}

	flight.FlightNo = input.FlightNo
	flight.Departure = input.Departure
	flight.Destination = input.Destination
	flight.DepartureTime = input.DepartureTime
	flight.ArrivalTime = input.ArrivalTime
	flight.FlightClass = input.FlightClass
	flight.Price = input.Price
	flight.RemainingCapacity = input.RemainingCapacity
	flight.CancelCondition = input.CancelCondition

	addNewFlight := db.Debug().Create(&flight).Commit()

	if addNewFlight.RowsAffected < 1 {
		return &flight, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &flight, nil
}

func (a FlightRepositoryImp) Update(input *domain.Flight) (*domain.Flight, error) {
	var flight domain.Flight
	_, err := a.Get(int(input.ID))
	if err != nil {
		return nil, err
	}
	db, _ := database.GetDatabaseConnection()

	println(input.RemainingCapacity)
	tx := db.Save(&input)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func (a FlightRepositoryImp) Get(id int) (*domain.Flight, error) {
	var flight domain.Flight
	db, _ := database.GetDatabaseConnection()
	tx := db.First(&flight, id)
	if tx.Error != nil {
		return &flight, tx.Error
	}
	return &flight, nil
}

func (a FlightRepositoryImp) GetAll() (*[]domain.Flight, error) {
	var flights []domain.Flight
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&flights)

	checkFlightExist := db.Debug().Find(&flights)

	if checkFlightExist.RowsAffected <= 0 {
		return &flights, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&flights)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &flights, nil
}

func (a FlightRepositoryImp) Delete(id int) error {
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
