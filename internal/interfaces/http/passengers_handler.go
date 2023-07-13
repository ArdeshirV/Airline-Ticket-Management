package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

func PassengerRoute(passengerHandler passengerHandler, e *echo.Echo) {
	e.POST("/passengers", passengerHandler.AddPassenger)
	e.GET("/passengers", passengerHandler.GetPassengers)
	e.PUT("/passengers/:id", passengerHandler.UpdatePassenger)
	e.DELETE("/passengers/:id", passengerHandler.DeletePassenger)
}

type passengerHandler struct {
	passengerRepo persistence.PassengerRepository
	userhandler   UserHandler
}

func NewPassegerHandler(passengerRepo persistence.PassengerRepository,
	userhandler UserHandler) passengerHandler {
	return passengerHandler{
		passengerRepo: passengerRepo,
		userhandler:   userhandler,
	}
}
func (p passengerHandler) AddPassenger(c echo.Context) error {
	user, err := p.userhandler.GetUserFromSession(c)
	println(":::::::::::::::::1")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Login first"})
	}
	println(user)
	println("UID:", user.ID)
	var passenger domain.Passenger
	err = c.Bind(&passenger)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	println(":::::::::::::::::2")
	passenger.UserID = user.ID
	pr := p.passengerRepo
	if _, err := pr.Create(&passenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}

	println(":::::::::::::::::3")
	msg := fmt.Sprintf("passenger with ID:%d added", passenger.ID)
	return echoStringAsJSON(c, http.StatusOK, msg)
}

func (p passengerHandler) DeletePassenger(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echoStringAsJSON(c, http.StatusBadRequest, "the parameter 'id' is required")
	}
	user, err := p.userhandler.GetUserFromSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Login first"})
	}
	println("UID:", user.ID)
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	pr := p.passengerRepo
	passenger, err := pr.Get(passengerID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if passenger.UserID != user.ID {
		return c.NoContent(http.StatusUnauthorized)
	}
	if err = pr.Delete(passengerID); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	msg := fmt.Sprintf("passenger with ID:%d deleted", passengerID)
	return echoStringAsJSON(c, http.StatusOK, msg)
}

func (p passengerHandler) UpdatePassenger(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echoStringAsJSON(c, http.StatusBadRequest, "the parameter 'id' is required")
	}
	user, err := p.userhandler.GetUserFromSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Login first"})
	}
	println("UID:", user.ID)
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	var updatedPassenger domain.Passenger
	err = c.Bind(&updatedPassenger)
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	println("UID:", user.ID)
	pr := p.passengerRepo
	passenger, err := pr.Get(passengerID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if passenger.UserID != user.ID {
		return c.NoContent(http.StatusUnauthorized)
	}
	if _, err := pr.Update(&updatedPassenger); err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	msg := fmt.Sprintf("passenger with ID:%d updated", passengerID)
	return echoStringAsJSON(c, http.StatusOK, msg)
}

func (p passengerHandler) GetPassengers(c echo.Context) error {
	user, err := p.userhandler.GetUserFromSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Login first"})
	}
	pr := p.passengerRepo
	response, err := pr.GetByUserId(int(user.ID))
	if err != nil {
		return echoErrorAsJSON(c, http.StatusBadRequest, err)
	}
	return echoJSON(c, http.StatusOK, response)
}
