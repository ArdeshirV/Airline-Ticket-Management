package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type AirportRepository struct {
}

func (a *AirportRepository) New() *AirportRepository {
	return &AirportRepository{}
}

func (a *AirportRepository) Create(input *domain.Airport) (*domain.Airport, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (a *AirportRepository) Update(input *domain.Airport) (*domain.Airport, error) {
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

func (a *AirportRepository) Get(id int) (*domain.Airport, error) {
	var airport domain.Airport
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airport)

	checkAirportExist := db.Debug().Where("ID = ?", id)

	tx := checkAirportExist.First(&airport)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &airport, nil
}

func (a *AirportRepository) GetAll() (*[]domain.Airport, error) {
	var airports []domain.Airport
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airports)

	tx := db.Debug().Find(&airports)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &airports, nil
}

func (a *AirportRepository) Delete(id int) error {
	airport, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&airport)
	deleted := db.Debug().Delete(airport).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
