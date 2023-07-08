package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

func PassengerRoute(e *echo.Echo) {
	e.POST("/passengers", AddPassenger)
	e.GET("/passengers/:id", GetPassenger)
	e.PUT("/passengers/:id", UpdatePassenger)
	e.DELETE("/passengers/:id", DeletePassenger)
}

func AddPassenger(c echo.Context) error {
	var newPassenger *domain.Passenger
	if err := c.Bind(newPassenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := persistence.NewPassengerRepository()
	response, err := pr.Create(newPassenger)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	return echoJSON(c, http.StatusOK, *response)
}

func DeletePassenger(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echoStringAsJSON(c, http.StatusBadRequest, "the parameter 'id' is required")
	}
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := persistence.NewPassengerRepository()
	if err = pr.Delete(passengerID); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	msg := fmt.Sprintf("passenger with ID:%d deleted", passengerID)
	return echoStringAsJSON(c, http.StatusOK, msg)
}

func UpdatePassenger(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echoStringAsJSON(c, http.StatusBadRequest, "the parameter 'id' is required")
	}
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	var updatedPassenger *domain.Passenger
	if err := c.Bind(updatedPassenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := persistence.NewPassengerRepository()
	if _, err := pr.Update(updatedPassenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	msg := fmt.Sprintf("passenger with ID:%d updated", passengerID)
	return echoStringAsJSON(c, http.StatusOK, msg)
}

func GetPassenger(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echoStringAsJSON(c, http.StatusBadRequest, "the parameter 'id' is required")
	}
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := persistence.NewPassengerRepository()
	response, err := pr.Get(passengerID)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	return echoJSON(c, http.StatusOK, response)
}
