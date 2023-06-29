package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/internal/usecase"
)

const (
	ParamTicketID = "ticketid"
)

func PrintTicketRoute(e *echo.Echo) {
	e.GET("/print_ticket", ticketHandler)
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
