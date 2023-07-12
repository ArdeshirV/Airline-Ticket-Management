package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type AirlineRepository struct {
	// todo: you could have a database connection as your repository struct, so you don't have to use something like this in each method: db, _ := database.GetDatabaseConnection()
}

func NewAirlineRepsoitory() *AirlineRepository {
	return &AirlineRepository{}
}

func (r *AirlineRepository) Create(input *domain.Airline) (*domain.Airline, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (r *AirlineRepository) Update(input *domain.Airline) (*domain.Airline, error) {
	db, _ := database.GetDatabaseConnection()
	_, err := r.Get(int(input.ID))
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

func (r *AirlineRepository) Get(id int) (*domain.Airline, error) {
	var airline domain.Airline
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airline)

	checkAirlineExist := db.Debug().Where("ID = ?", id)

	tx := checkAirlineExist.First(&airline)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &airline, nil
}

func (r *AirlineRepository) GetAll() (*[]domain.Airline, error) {
	var airlines []domain.Airline
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airlines)

	tx := db.Debug().Find(&airlines)

	if err := tx.Error; err != nil {
		return nil, err
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
	deleted := db.Debug().Delete(airline)
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
