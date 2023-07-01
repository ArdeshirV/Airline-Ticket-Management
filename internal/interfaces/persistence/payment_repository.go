package persistence

import (
	"errors"
	"net/http"
	"strconv"

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
	var payment domain.Payment
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&payment)

	checkPaymentExist := db.Debug().Where(&payment, "ID = ?", input.ID)

	if checkPaymentExist.RowsAffected <= 0 {
		return &payment, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkPaymentExist.Update("ID", input.ID).Update("PayAmount", input.PayAmount).Update("PaymentSerial", input.PaymentSerial).Update("PayTime", input.PayTime)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedPayment := tx.Commit()
		if updatedPayment.RowsAffected < 1 {
			return &payment, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &payment, nil
}

func (a PaymentRepositoryImp) Get(id int) (*domain.Payment, error) {
	var payment domain.Payment
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&payment)

	checkPaymentExist := db.Debug().Where(&payment, "ID = ?", id)

	if checkPaymentExist.RowsAffected <= 0 {
		return &payment, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkPaymentExist.Find(&payment)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &payment, nil
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
