package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
)

const (
	ParamTicketID = "ticketid"
)

func PrintTicketRoute(e *echo.Echo) {
	e.GET("/print_ticket", PrintTicketHandler)
}

func PrintTicketHandler(ctx echo.Context) error {
	ticketid := ctx.QueryParam(ParamTicketID)
	if ticketid == "" {
		errMsg := "the 'Ticketid' parameter is required"
		return echoStringAsJSON(ctx, http.StatusBadRequest, errMsg)
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
