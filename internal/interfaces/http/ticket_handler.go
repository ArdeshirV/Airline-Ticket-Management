package http

import (
	"fmt"
	"github.com/the-go-dragons/final-project/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
)

const (
	ParamTicketID = "ticketid"
)

type TicketActionRequest struct {
	UserId   int
	TicketId int
}

type tickets *[]domain.Ticket

type TicketHandler struct {
	ticketUseCase *usecase.TicketUsecase
	flightUseCase *usecase.FlightUseCase
	booking       *usecase.Booking
}

func NewTicketHandler(ticketUseCase *usecase.TicketUsecase, flightUseCase *usecase.FlightUseCase, booking *usecase.Booking) *TicketHandler {
	return &TicketHandler{
		ticketUseCase: ticketUseCase,
		flightUseCase: flightUseCase,
		booking:       booking,
	}
}

func TicketRoute(e *echo.Echo) {
	e.GET("/ticket", ticketHandler)
}

func ticketHandler(ctx echo.Context) error {
	ticketid := ctx.QueryParam(ParamTicketID)
	if ticketid == "" {
		return echoStringAsJSON(ctx, http.StatusBadRequest, "the 'Ticketid' parameter is required")
	}
	id, err := strconv.Atoi(ticketid)
	if err != nil {
		errMsg := fmt.Errorf("failed to convert ticketid='%v' to integer", ticketid)
		return echoErrorAsJSON(ctx, http.StatusBadRequest, errMsg)
	}
	err = usecase.CreateTicketAsPDF(int(id), config.Config.App.TicketFileName)
	if err != nil {
		errMsg := fmt.Errorf("%v", err)
		return echoErrorAsJSON(ctx, http.StatusBadRequest, errMsg)
	}
	return ctx.File(config.Config.App.TicketFileName)
}

func (th *TicketHandler) Cancel(c echo.Context) error {
	var request TicketActionRequest

	// Check the body data
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, MassageResponse{Message: "Invalid body request"})

	}
	if strconv.Itoa(request.TicketId) == "" {
		return c.JSON(http.StatusBadRequest, MassageResponse{Message: "Missing required fields"})
	}

	ticket, err := th.booking.CancelTicket(request.TicketId)
	if err == nil {
		return c.JSON(http.StatusConflict, MassageResponse{Message: "cancellation failed"})
	}

	err = th.flightUseCase.IncreaseFlightCapacity(&ticket.Flight)
	if err == nil {
		return c.JSON(http.StatusConflict, MassageResponse{Message: "cancellation failed"})
	}

	return c.JSON(http.StatusOK, MassageResponse{Message: "Cancelled Successfully"})
}

func (th *TicketHandler) GetReservedUsers(c echo.Context) error {

	// Check the body data
	userId := c.QueryParam("userId")
	println(userId)
	if userId == "0" || userId == "" {
		return c.JSON(http.StatusBadRequest, MassageResponse{Message: "Invalid body request"})

	}

	atoi, err := strconv.Atoi(userId)
	if err != nil {
		return c.JSON(http.StatusConflict, MassageResponse{Message: "can not read the ticket data"})
	}
	tickets, err := th.booking.GetReservedTickets(uint(atoi))
	if err == nil {
		return c.JSON(http.StatusConflict, MassageResponse{Message: "can not read the ticket data"})
	}

	return c.JSON(http.StatusOK, tickets)
}
