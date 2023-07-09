package persistence

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type TicketRepository struct {
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (a *TicketRepository) Create(input *domain.Ticket) (*domain.Ticket, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (a *TicketRepository) Update(input *domain.Ticket) (*domain.Ticket, error) {
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

func (a *TicketRepository) Get(id int) (*domain.Ticket, error) {
	var ticket domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&ticket)

	checkTicketExist := db.Debug().Where("ID = ?", id)

	tx := checkTicketExist.First(&ticket)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (a *TicketRepository) GetAll() (*[]domain.Ticket, error) {
	var tickets []domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&tickets)

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

func (a *TicketRepository) GetAllNotArrivedByUserId() (*[]domain.Ticket, error) {
	var tickets []domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&tickets)

	err := db.Table("tickets").
		Select("*").InnerJoins("inner join flights on tickets.FlightID = flights.ID").
		Where("Refund = ?", false).
		Where("flights.DepartureTime > ?", time.Now()).
		Scan(&tickets).Error

	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	return &tickets, nil
}

func (a *TicketRepository) CreateList(input *[]domain.Ticket) (*[]domain.Ticket, error) {
	db, _ := database.GetDatabaseConnection()

	tx := db.Debug().Create(&input)

	if tx.Error != nil {
		return input, tx.Error
	}

	return input, nil
}

func (a *TicketRepository) GetByUserId(userId uint) (*[]domain.Ticket, error) {
	var tickets []domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&tickets)

	tx := db.
		Table("tickets").
		Joins("join flights on tickets.flight_id = flights.id").
		Where("user_id = ?", userId).
		Where("Refund = ?", false).
		Where("flights.departure_time > ?", time.Now()).
		Preload("Passenger").
		Preload("Flight.Departure").
		Preload("Flight.Destination").
		Find(&tickets)

	if tx.Error != nil {
		return nil, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	return &tickets, nil
}

func (a *TicketRepository) GetCancelledByUserId(userId uint) (*[]domain.Ticket, error) {
	var tickets []domain.Ticket
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&tickets)

	tx := db.
		Table("tickets").
		Joins("join flights on tickets.Flight_ID = flights.ID").
		Where("user_ID = ?", userId).
		Where("Refund = ?", true).
		Where("flights.Departure_Time > ?", time.Now()).
		Find(&tickets)

	if tx.Error != nil {
		return nil, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	return &tickets, nil
}

func (a *TicketRepository) CancelTicket(id int) (*domain.Ticket, error) {
	ticket, err := a.Get(id)

	if err != nil {
		return nil, err
	}

	ticket.Refund = true

	ticket, err = a.Update(ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
