package persistence

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type TicketRepository struct {
}

func (a *TicketRepository) New() *TicketRepository {
	return &TicketRepository{}
}

func (a *TicketRepository) Create(input *domain.Ticket) (*domain.Ticket, error) {
	var ticket domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&ticket)

	checkTicketExist := db.Debug().First(&ticket, "ID = ?", input.ID)

	if checkTicketExist.RowsAffected > 0 {
		return &ticket, errors.New(strconv.Itoa(http.StatusConflict))
	}

	ticket.Flight = input.Flight
	ticket.Passenger = input.Passenger
	ticket.Payment = input.Payment
	ticket.User = input.User
	ticket.Refund = input.Refund

	addNewTicket := db.Debug().Create(&ticket).Commit()

	if addNewTicket.RowsAffected < 1 {
		return &ticket, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &ticket, nil
}

func (a *TicketRepository) Update(input *domain.Ticket) (*domain.Ticket, error) {
	var ticket domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&ticket)

	checkTicketExist := db.Debug().Where(&ticket, "ID = ?", input.ID)

	if checkTicketExist.RowsAffected <= 0 {
		return &ticket, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkTicketExist.Update("ID", input.ID).Update("Flight", input.Flight).Update("Passenger", input.Passenger)
	tx = tx.Update("Payment", input.Payment).Update("User", input.User).Update("Refund", input.Refund)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updateTicket := tx.Commit()
		if updateTicket.RowsAffected < 1 {
			return &ticket, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &ticket, nil
}

func (a *TicketRepository) Get(id int) (*domain.Ticket, error) {
	var ticket domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&ticket)

	checkTicketExist := db.Debug().Where(&ticket, "ID = ?", id)

	if checkTicketExist.RowsAffected <= 0 {
		return &ticket, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkTicketExist.Find(&ticket)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (a *TicketRepository) GetAll() (*[]domain.Ticket, error) {
	var tickets []domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&tickets)

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

func (a *TicketRepository) Delete(id int) error {
	ticket, err := a.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&ticket)
	deleted := db.Debug().Delete(ticket).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
