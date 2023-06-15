package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type FlightRepo struct {
	db *gorm.DB
}

func (a *FlightRepo) New(db *gorm.DB) *FlightRepo {
	return &FlightRepo{db: db}
}

func (a *FlightRepo) Save(input *domain.Flight) (*domain.Flight, error) {
	var flight domain.Flight
	db := a.db.Model(&flight)

	checkFlightExist := db.Debug().First(&flight, "ID = ?", input.ID)

	if checkFlightExist.RowsAffected > 0 {
		return &flight, errors.New(strconv.Itoa(http.StatusConflict))
	}

	flight.FlightNo = input.FlightNo
	flight.Departure = input.Departure
	flight.Destination = input.Destination
	flight.DepartureTime = input.DepartureTime
	flight.ArrivalTime = input.ArrivalTime
	flight.Airlines = input.Airlines
	flight.FlightClass = input.FlightClass
	flight.Price = input.Price
	flight.Capacity = input.Capacity
	flight.CancelCondition = input.CancelCondition

	addNewFlight := db.Debug().Create(&flight).Commit()

	if addNewFlight.RowsAffected < 1 {
		return &flight, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &flight, nil
}

func (a *FlightRepo) Update(input *domain.Flight) (*domain.Flight, error) {
	var flight domain.Flight
	db := a.db.Model(&flight)

	checkCityExist := db.Debug().Where(&flight, "ID = ?", input.ID)

	if checkCityExist.RowsAffected <= 0 {
		return &flight, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkCityExist.Update("ID", input.ID).Update("FlightNo", input.FlightNo).Update("Departure", input.Departure)
	tx = tx.Update("Destination", input.Destination).Update("DepartureTime", input.DepartureTime).Update("ArrivalTime", input.ArrivalTime)
	tx = tx.Update("Airlines", input.Airlines).Update("FlightClass", input.FlightClass).Update("Price", input.Price)
	tx = tx.Update("Capacity", input.Capacity).Update("CancelCondition", input.CancelCondition)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedFlight := tx.Commit()
		if updatedFlight.RowsAffected < 1 {
			return &flight, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &flight, nil
}

func (a *FlightRepo) Get(id int) (*domain.Flight, error) {
	var flight domain.Flight
	db := a.db.Model(&flight)

	checkFlightExist := db.Debug().Where(&flight, "ID = ?", id)

	if checkFlightExist.RowsAffected <= 0 {
		return &flight, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkFlightExist.Find(&flight)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func (a *FlightRepo) GetAll() (*[]domain.Flight, error) {
	var flights []domain.Flight
	db := a.db.Model(&flights)

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

func (a *FlightRepo) Delete(id int) error {
	flight, err := a.Get(id)
	if err != nil {
		return err
	}
	db := a.db.Model(&flight)
	deleted := db.Debug().Delete(flight).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
