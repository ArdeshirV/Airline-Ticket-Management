package persistence

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type AirlineRepository struct {
}

func (r *AirlineRepository) New() *AirlineRepository {
	return &AirlineRepository{}
}

func (r *AirlineRepository) Create(input *domain.Airline) (*domain.Airline, error) {
	var airline domain.Airline
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airline)

	checkAirlineExist := db.Debug().First(&airline, "ID = ?", input.ID)

	if checkAirlineExist.RowsAffected > 0 {
		return &airline, errors.New(strconv.Itoa(http.StatusConflict))
	}

	airline.Name = input.Name
	airline.Logo = input.Logo

	addNewAirline := db.Debug().Create(&airline).Commit()

	if addNewAirline.RowsAffected < 1 {
		return &airline, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &airline, nil
}

func (r *AirlineRepository) Update(input *domain.Airline) (*domain.Airline, error) {
	var airline domain.Airline
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airline)

	checkAirlineExist := db.Debug().Where(&airline, "ID = ?", input.ID)

	if checkAirlineExist.RowsAffected <= 0 {
		return &airline, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkAirlineExist.Update("ID", input.ID).Update("Name", input.Name).Update("Logo", input.Logo)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirline := tx.Commit()
		if updatedAirline.RowsAffected < 1 {
			return &airline, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airline, nil
}

func (r *AirlineRepository) Get(id int) (*domain.Airline, error) {
	var airline domain.Airline
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airline)

	checkAirlineExist := db.Debug().Where(&airline, "ID = ?", id)

	if checkAirlineExist.RowsAffected <= 0 {
		return &airline, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Where("ID = ?", id).Find(&airline)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirline := tx.Commit()
		if updatedAirline.RowsAffected < 1 {
			return &airline, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airline, nil
}

func (r *AirlineRepository) GetAll() (*[]domain.Airline, error) {
	var airlines []domain.Airline
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airlines)

	checkAirlineExist := db.Debug().Find(&airlines)

	if checkAirlineExist.RowsAffected <= 0 {
		return &airlines, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&airlines)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirline := tx.Commit()
		if updatedAirline.RowsAffected < 1 {
			return &airlines, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airlines, nil
}

func (r *AirlineRepository) Delete(id int) error {
	airline, err := r.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airline)
	deleted := db.Debug().Delete(airline).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
