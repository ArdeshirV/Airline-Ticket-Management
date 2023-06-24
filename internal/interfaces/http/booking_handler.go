package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
)

type BookingRequest struct {
	FlightID     int
	PassengerIDs []int
}

type BookingError struct {
	Message string
}

type BookingResponse struct {
	OrderID uint
}

type FainalizeRequest struct {
	OrderID int
}
type BookingHandler struct {
	booking *usecase.Booking
}

func NewBookingHandler(booking *usecase.Booking) *BookingHandler {
	return &BookingHandler{booking: booking}
}

func (b *BookingHandler) Book(c echo.Context) error {
	var request BookingRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BookingError{Message: "Invalid body request"})
	}

	if request.FlightID == 0 || len(request.PassengerIDs) == 0 {
		return c.JSON(http.StatusBadRequest, BookingError{Message: "Missing required fields"})
	}

	orderID, err := b.booking.Book(request.FlightID, request.PassengerIDs)
	if err != nil {
		switch err.(type) {
		case usecase.FlightNotFound:
			return c.JSON(http.StatusNotFound, BookingError{Message: err.Error()})
		case usecase.FlightCapacityError:
			return c.JSON(http.StatusBadRequest, BookingError{Message: err.Error()})
		case usecase.SomePassengerNotFound:
			return c.JSON(http.StatusNotFound, BookingError{Message: err.Error()})
		}
	}
	return c.JSON(http.StatusAccepted, BookingResponse{orderID})
}

func (b *BookingHandler) Finalize(c echo.Context) error {
	var request FainalizeRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BookingError{Message: "Invalid body request"})
	}

	if request.OrderID == 0 {
		return c.JSON(http.StatusBadRequest, BookingError{Message: "Missing required fields"})
	}

	err = b.booking.Finalize(request.OrderID)
	if err != nil {
		switch err.(type) {
		case usecase.OrderNotFound:
			return c.JSON(http.StatusNotFound, BookingError{Message: err.Error()})
		case usecase.OrderNotPaid:
			return c.JSON(http.StatusBadRequest, BookingError{Message: err.Error()})
		case usecase.OrderItemsNotFound:
			return c.JSON(http.StatusNotFound, BookingError{Message: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, BookingError{Message: err.Error()})

		}
	}
	return c.NoContent(http.StatusAccepted)
}
