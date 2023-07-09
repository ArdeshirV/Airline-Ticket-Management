package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type PaymentRepository interface {
	Create(input *domain.Payment) (*domain.Payment, error)
	Update(input *domain.Payment) (*domain.Payment, error)
	Get(id int) (*domain.Payment, error)
	GetByOrderId(orderID int) (*domain.Payment, error)
}
type PaymentRepositoryImp struct {
}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImp{}
}

func (a PaymentRepositoryImp) Create(input *domain.Payment) (*domain.Payment, error) {
	db, _ := database.GetDatabaseConnection()
	tx := db.Debug().Create(&input)

	if tx.Error != nil {
		return input, tx.Error
	}

	return input, nil
}

func (a PaymentRepositoryImp) Update(input *domain.Payment) (*domain.Payment, error) {
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

func (a PaymentRepositoryImp) Get(id int) (*domain.Payment, error) {
	var payment domain.Payment
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&payment)

	checkTicketExist := db.Debug().Where("ID = ?", id)

	tx := checkTicketExist.First(&payment)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (a PaymentRepositoryImp) GetAll() (*[]domain.Payment, error) {
	var payments []domain.Payment
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&payments)

	tx := db.Debug().Find(&payments)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &payments, nil
}

func (a PaymentRepositoryImp) Delete(id int) error {
	payment, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&payment)
	deleted := db.Debug().Delete(payment).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (a PaymentRepositoryImp) GetByOrderId(orderID int) (*domain.Payment, error) {
	var payment domain.Payment
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&payment)

	tx := db.Debug().Where("order_id = ?", orderID).First(&payment)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &payment, nil
}
