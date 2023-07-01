package persistence

import (
	"errors"
	"net/http"
	"strconv"

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
}
type PassengerRepositoryImp struct {
}

func NewPassengerRepository() PassengerRepository {
	return &PassengerRepositoryImp{}
}

func (a PassengerRepositoryImp) Create(input *domain.Passenger) (*domain.Passenger, error) {
	var passenger domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passenger)

	checkPassengerExist := db.Debug().First(&passenger, "ID = ?", input.ID)

	if checkPassengerExist.RowsAffected > 0 {
		return &passenger, errors.New(strconv.Itoa(http.StatusConflict))
	}

	passenger.FirstName = input.FirstName
	passenger.LastName = input.LastName
	passenger.NationalCode = input.NationalCode
	passenger.Email = input.Email
	passenger.Gender = input.Gender
	passenger.Phone = input.Phone
	passenger.Address = input.Address
	passenger.Tickets = input.Tickets

	addNewPassenger := db.Debug().Create(&passenger).Commit()

	if addNewPassenger.RowsAffected < 1 {
		return &passenger, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &passenger, nil
}

func (a PassengerRepositoryImp) Update(input *domain.Passenger) (*domain.Passenger, error) {
	var passenger domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passenger)

	checkPassengerExist := db.Debug().Where(&passenger, "ID = ?", input.ID)

	if checkPassengerExist.RowsAffected <= 0 {
		return &passenger, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkPassengerExist.Update("ID", input.ID).Update("FirstName", input.FirstName).Update("LastName", input.LastName)
	tx = tx.Update("NationalCode", input.NationalCode).Update("Email", input.Email).Update("Gender", input.Gender)
	tx = tx.Update("Phone", input.Phone).Update("Address", input.Address).Update("Tickets", input.Tickets)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedPassenger := tx.Commit()
		if updatedPassenger.RowsAffected < 1 {
			return &passenger, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &passenger, nil
}

func (a PassengerRepositoryImp) Get(id int) (*domain.Passenger, error) {
	var passenger domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passenger)

	checkPassengerExist := db.Debug().Where(&passenger, "ID = ?", id)

	if checkPassengerExist.RowsAffected <= 0 {
		return &passenger, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkPassengerExist.Find(&passenger)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &passenger, nil
}

func (a PassengerRepositoryImp) GetAll() (*[]domain.Passenger, error) {
	var passengers []domain.Passenger
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&passengers)

	checkPassengerExist := db.Debug().Find(&passengers)

	if checkPassengerExist.RowsAffected <= 0 {
		return &passengers, errors.New(strconv.Itoa(http.StatusNotFound))
	}

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
	db = db.Model(&passenger)
	deleted := db.Debug().Delete(passenger).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
