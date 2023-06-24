package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
)

const (
	ParamTicketID = "ticketid"
	TicketFileName = "pdf/ticket.pdf"  // TODO: Put it into env & config file
)

func TicketRoute(e *echo.Echo) {
	e.GET("/ticket", ticketHandler)
}

func ticketHandler(ctx echo.Context) error {
	ticketid := ctx.QueryParam(ParamTicketID)
	if ticketid == "" {
		return echoStringAsJSON(ctx, http.StatusBadRequest, "the 'Ticketid' parameter is required")
	}
	id, err := strconv.ParseInt(ticketid, 10, 64)
	if err != nil {
		errMsg := fmt.Errorf("failed to convert ticketid='%v' to integer", ticketid)
		return echoErrorAsJSON(ctx, http.StatusBadRequest, errMsg)
	}
	err = usecase.CreateTicketAsPDF(int(id), TicketFileName)
	if err != nil {
		errMsg := fmt.Errorf("%v", err)
		return echoErrorAsJSON(ctx, http.StatusBadRequest, errMsg)
	}
	return ctx.File(TicketFileName)
}
