package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type TicketRepo struct {
	db *gorm.DB
}

func (a *TicketRepo) New(db *gorm.DB) *TicketRepo {
	return &TicketRepo{db: db}
}

func (a *TicketRepo) Save(input *domain.Ticket) (*domain.Ticket, error) {
	var ticket domain.Ticket
	db := a.db.Model(&ticket)

	checkTicketExist := db.Debug().First(&ticket, "ID = ?", input.ID)

	if checkTicketExist.RowsAffected > 0 {
		return &ticket, errors.New(strconv.Itoa(http.StatusConflict))
	}

	ticket.Flight = input.Flight
	ticket.Passenger = input.Passenger
	ticket.Payment = input.Payment
	ticket.User = input.User
	ticket.StatusPay = input.StatusPay
	ticket.Refund = input.Refund

	addNewTicket := db.Debug().Create(&ticket).Commit()

	if addNewTicket.RowsAffected < 1 {
		return &ticket, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &ticket, nil
}

func (a *TicketRepo) Update(input *domain.Ticket) (*domain.Ticket, error) {
	var city domain.Ticket
	db := a.db.Model(&city)

	checkTicketExist := db.Debug().Where(&city, "ID = ?", input.ID)

	if checkTicketExist.RowsAffected <= 0 {
		return &city, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkTicketExist.Update("ID", input.ID).Update("Flight", input.Flight).Update("Passenger", input.Passenger)
	tx = tx.Update("Payment", input.Payment).Update("User", input.User).Update("StatusPay", input.StatusPay)
	tx = tx.Update("Refund", input.Refund)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updateTicket := tx.Commit()
		if updateTicket.RowsAffected < 1 {
			return &city, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &city, nil
}

func (a *TicketRepo) Get(id int) (*domain.Ticket, error) {
	var city domain.Ticket
	db := a.db.Model(&city)

	checkTicketExist := db.Debug().Where(&city, "ID = ?", id)

	if checkTicketExist.RowsAffected <= 0 {
		return &city, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkTicketExist.Find(&city)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (a *TicketRepo) GetAll() (*[]domain.Ticket, error) {
	var tickets []domain.Ticket
	db := a.db.Model(&tickets)

	checkTicketExist := db.Debug().Find(&tickets)

	if checkTicketExist.RowsAffected <= 0 {
		return &tickets, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&tickets)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &tickets, nil
}

func (a *TicketRepo) Delete(id int) error {
	city, err := a.Get(id)
	if err != nil {
		return err
	}
	db := a.db.Model(&city)
	deleted := db.Debug().Delete(city).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
