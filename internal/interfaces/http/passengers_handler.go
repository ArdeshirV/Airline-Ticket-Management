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
	var passenger domain.Passenger
	err := c.Bind(&passenger)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := persistence.NewPassengerRepository()
	if _, err := pr.Create(&passenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	msg := fmt.Sprintf("passenger with ID:%d added", passenger.ID)
	return echoStringAsJSON(c, http.StatusOK, msg)
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
	var updatedPassenger domain.Passenger
	err = c.Bind(&updatedPassenger)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := persistence.NewPassengerRepository()
	if _, err := pr.Update(&updatedPassenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	msg := fmt.Sprintf("passenger with ID:%d updated", passengerID)
	println("Hey:", msg)
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
