package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type CityRepository struct {
}

func NewCityRepository() *CityRepository {
	return &CityRepository{}
}

func (a *CityRepository) Create(input *domain.City) (*domain.City, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (a *CityRepository) Update(input *domain.City) (*domain.City, error) {
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

func (a *CityRepository) Get(id int) (*domain.City, error) {
	var city domain.City
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&city)

	checkCityExist := db.Debug().Where("ID = ?", id)

	tx := checkCityExist.First(&city)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (a *CityRepository) GetAll() (*[]domain.City, error) {
	var cities []domain.City
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&cities)

	tx := db.Debug().Find(&cities)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &cities, nil
}

func (a *CityRepository) Delete(id int) error {
	city, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&city)
	deleted := db.Debug().Delete(city)
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
