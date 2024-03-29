package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type PassengerRepository interface {
	Create(input *domain.Passenger) (*domain.Passenger, error)
	Update(input *domain.Passenger) (*domain.Passenger, error)
	Get(id int) (*domain.Passenger, error)
	GetAll() (*[]domain.Passenger, error)
	GetList(IDs []int) ([]domain.Passenger, error)
	Delete(id int) error
	GetByUserId(id int) (*[]domain.Passenger, error)
}

type PassengerRepositoryImp struct {
}

func NewPassengerRepository() PassengerRepository {
	return &PassengerRepositoryImp{}
}

func (a PassengerRepositoryImp) Create(input *domain.Passenger) (*domain.Passenger, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (a PassengerRepositoryImp) Update(input *domain.Passenger) (*domain.Passenger, error) {
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

func (a PassengerRepositoryImp) Get(id int) (*domain.Passenger, error) {
	var passenger domain.Passenger
	db, _ := database.GetDatabaseConnection()

	tx := db.First(&passenger, id)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &passenger, nil
}

func (a PassengerRepositoryImp) GetByUserId(id int) (*[]domain.Passenger, error) {
	var passenger []domain.Passenger
	db, _ := database.GetDatabaseConnection()

	tx := db.Where("user_id = ?", id).Find(&passenger)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &passenger, nil
}

func (a PassengerRepositoryImp) GetAll() (*[]domain.Passenger, error) {
	var passengers []domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passengers)

	tx := db.Debug().Find(&passengers)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &passengers, nil
}

func (a PassengerRepositoryImp) GetList(IDs []int) ([]domain.Passenger, error) {
	var passengers []domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passengers)
	db.Find(&passengers, IDs)
	return passengers, nil
}

func (a PassengerRepositoryImp) Delete(id int) error {
	passenger, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	deleted := db.Debug().Delete(passenger)
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
