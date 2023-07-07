package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type AirplaneRepository interface {
	Create(input *domain.Airplane) (*domain.Airplane, error)
	Update(input *domain.Airplane) (*domain.Airplane, error)
	Get(id int) (*domain.Airplane, error)
	Delete(id int) error
}

type AirplaneRepositoryImpl struct {
	// todo: you could have a database connection as your repository struct, so you don't have to use something like this in each method: db, _ := database.GetDatabaseConnection()
}

func NewAirplaneRepository() AirplaneRepository {
	return &AirplaneRepositoryImpl{}
}

func (r AirplaneRepositoryImpl) Create(input *domain.Airplane) (*domain.Airplane, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	tx := db.Create(&input)
	if tx.Error != nil {
		return input, tx.Error
	}
	return input, nil
}

func (r AirplaneRepositoryImpl) Update(input *domain.Airplane) (*domain.Airplane, error) {
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

func (r AirplaneRepositoryImpl) Get(id int) (*domain.Airplane, error) {
	var airplane domain.Airplane
	db, _ := database.GetDatabaseConnection()

	tx := db.Debug().Where(&airplane, "ID = ?", id).First(&airplane)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &airplane, nil
}

func (r AirplaneRepositoryImpl) Delete(id int) error {
	airplane, err := r.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	deleted := db.Debug().Delete(airplane).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
