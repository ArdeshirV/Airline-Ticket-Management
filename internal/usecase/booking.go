package usecase

import (
	"github.com/google/uuid"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type Booking struct {
	flightRepo    persistence.FlightRepository
	passengerRepo persistence.PassengerRepository
	orderRepo     persistence.OrderRepository
	paymentRepo   persistence.PaymentRepository
	ticketRepo    *persistence.TicketRepository
}

func NewBooking(flightRepo persistence.FlightRepository,
	passengerRepo persistence.PassengerRepository,
	orderRepo persistence.OrderRepository,
	paymentRepo persistence.PaymentRepository) *Booking {
	return &Booking{
		flightRepo:    flightRepo,
		passengerRepo: passengerRepo,
		orderRepo:     orderRepo,
		paymentRepo:   paymentRepo,
	}
}

func (b Booking) Book(flightID int, passengerIDs []int) (uint, error) {
	flight, err := b.flightRepo.Get(flightID)
	if err != nil {
		return 0, FlightNotFound{flightID}
	}
	if flight.RemainingCapacity < len(passengerIDs) {
		return 0, FlightCapacityError{flight.RemainingCapacity, len(passengerIDs)}
	}
	passengers, _ := b.passengerRepo.GetList(passengerIDs)
	if len(passengers) != len(passengerIDs) {
		return 0, SomePassengerNotFound{len(passengers), len(passengerIDs)}
	}
	var orderItems []domain.OrderItem
	for _, passenger := range passengers {
		item := domain.OrderItem{
			PassengerID: passenger.ID,
		}
		orderItems = append(orderItems, item)
	}
	order := domain.Order{
		OrderItems: orderItems,
		Amount:     flight.Price * int64(len(passengers)),
		FlightID:   flight.ID,
		Status:     domain.PENDING,
		OrderNum:   uuid.New().String(),
	}

	newOrder, err := b.orderRepo.Create(&order)
	if err != nil {
		return 0, err
	}

	return newOrder.ID, nil
}

func (b Booking) Finalize(orderID int) error {
	order, err := b.orderRepo.Get(orderID)
	if err != nil {
		println(orderID, err.Error())
		return OrderNotFound{orderID}
	}

	if order.Status == domain.PENDING {
		return OrderNotPaid{orderID}
	}

	if order.Status == domain.DELIVERED {
		return OrdrAlreadyDelivered{orderID}
	}

	flight, _ := b.flightRepo.Get(int(order.FlightID))
	items, err := b.orderRepo.GetItems(order.ID)
	if err != nil {
		return &OrderItemsNotFound{order.ID}
	}
	flight.RemainingCapacity = flight.RemainingCapacity - len(items)
	_, err = b.flightRepo.Update(flight)
	if err != nil {
		return err
	}
	payment, _ := b.paymentRepo.GetByOrderId(int(order.ID))
	var tickets []domain.Ticket
	for _, item := range items {
		ticket := domain.Ticket{
			FlightID:      order.FlightID,
			PassengerID:   item.PassengerID,
			PaymentID:     payment.ID,
			PaymentStatus: "PAID",
			Refund:        false,
			// Todo: must read user from security context
			UserID: 1,
		}
		tickets = append(tickets, ticket)
	}
	_, err = b.ticketRepo.CreateList(&tickets)
	if err != nil {
		return err
	}
	order.Status = domain.DELIVERED
	_, err = b.orderRepo.Update(order)
	if err != nil {
		return err
	}
	return nil
}
