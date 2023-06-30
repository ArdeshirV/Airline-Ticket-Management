package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type PassengerRepository struct {
}

func NewPassengerRepository() *PassengerRepository {
	return &PassengerRepository{}
}

func (a *PassengerRepository) Create(input *domain.Passenger) (*domain.Passenger, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (a *PassengerRepository) Update(input *domain.Passenger) (*domain.Passenger, error) {
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

func (a *PassengerRepository) Get(id int) (*domain.Passenger, error) {
	var passenger domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passenger)

	checkPassengerExist := db.Debug().Where(&passenger, "ID = ?", id)

	tx := checkPassengerExist.First(&passenger)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &passenger, nil
}

func (a *PassengerRepository) GetAll() (*[]domain.Passenger, error) {
	var passengers []domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passengers)

	tx := db.Debug().Find(&passengers)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &passengers, nil
}

func (a *PassengerRepository) GetList(IDs []int) ([]domain.Passenger, error) {
	var passengers []domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passengers)
	db.Find(&passengers, IDs)
	return passengers, nil
}

func (a *PassengerRepository) Delete(id int) error {
	passenger, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passenger)
	deleted := db.Debug().Delete(passenger).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
