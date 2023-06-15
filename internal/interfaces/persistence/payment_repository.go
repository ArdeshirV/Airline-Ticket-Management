package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type PaymentRepo struct {
	db *gorm.DB
}

func (a *PaymentRepo) New(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (a *PaymentRepo) Save(input *domain.Payment) (*domain.Payment, error) {
	var payment domain.Payment
	db := a.db.Model(&payment)

	checkPaymentExist := db.Debug().First(&payment, "ID = ?", input.ID)

	if checkPaymentExist.RowsAffected > 0 {
		return &payment, errors.New(strconv.Itoa(http.StatusConflict))
	}

	payment.PayTime = input.PayTime
	payment.PayAmount = input.PayAmount
	payment.PaymentSerial = input.PaymentSerial

	addNewPayment := db.Debug().Create(&payment).Commit()

	if addNewPayment.RowsAffected < 1 {
		return &payment, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &payment, nil
}

func (a *PaymentRepo) Update(input *domain.Payment) (*domain.Payment, error) {
	var payment domain.Payment
	db := a.db.Model(&payment)

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

func (a *PaymentRepo) Get(id int) (*domain.Payment, error) {
	var payment domain.Payment
	db := a.db.Model(&payment)

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

func (a *PaymentRepo) GetAll() (*[]domain.Payment, error) {
	var payments []domain.Payment
	db := a.db.Model(&payments)

	checkPaymentExist := db.Debug().Find(&payments)

	if checkPaymentExist.RowsAffected <= 0 {
		return &payments, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&payments)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &payments, nil
}

func (a *PaymentRepo) Delete(id int) error {
	payment, err := a.Get(id)
	if err != nil {
		return err
	}
	db := a.db.Model(&payment)
	deleted := db.Debug().Delete(payment).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
