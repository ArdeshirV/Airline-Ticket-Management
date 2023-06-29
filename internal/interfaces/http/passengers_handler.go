package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
)

var lastPassengerID uint
var passengers []domain.Passenger

func PassengerRoute(e *echo.Echo) {
	e.POST("/passengers", AddPassenger)
	e.DELETE("/passengers/:id", DeletePassenger)
	e.PUT("/passengers/:id", UpdatePassenger)
	e.GET("/passengers/:id", GetPassenger)
}

func AddPassenger(c echo.Context) error {
	var newPassenger domain.Passenger
	if err := c.Bind(&newPassenger); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	lastPassengerID++
	newPassenger.ID = lastPassengerID
	passengers = append(passengers, newPassenger)
	response := map[string]interface{}{
		"message":     "Passenger added successfully",
		"passengerId": newPassenger.ID,
	}
	return c.JSON(http.StatusOK, response)
}

func DeletePassenger(c echo.Context) error {
	passengerIDInt, err := strconv.Atoi(c.Param("id"))
	passengerID := uint(passengerIDInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid passenger ID"})
	}
	index := -1
	for i, p := range passengers {
		if p.ID == passengerID {
			index = i
			break
		}
	}
	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Passenger not found"})
	}
	passengers = append(passengers[:index], passengers[index+1:]...)
	response := map[string]string{
		"message": "Passenger deleted successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func UpdatePassenger(c echo.Context) error {
	passengerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid passenger ID"})
	}
	index := -1
	for i, p := range passengers {
		if p.ID == uint(passengerID) {
			index = i
			break
		}
	}
	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Passenger not found"})
	}
	var updatedPassenger domain.Passenger
	if err := c.Bind(&updatedPassenger); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	passengers[index].FirstName = updatedPassenger.FirstName
	passengers[index].LastName = updatedPassenger.LastName
	passengers[index].NationalCode = updatedPassenger.NationalCode
	passengers[index].Email = updatedPassenger.Email
	passengers[index].Gender = updatedPassenger.Gender
	passengers[index].BirthDate = updatedPassenger.BirthDate
	passengers[index].Phone = updatedPassenger.Phone
	passengers[index].Address = updatedPassenger.Address
	response := map[string]string{
		"message": "Passenger updated successfully",
	}
	return c.JSON(http.StatusOK, response)
}

func GetPassenger(c echo.Context) error {
	passengerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid passenger ID"})
	}
	var foundPassenger *domain.Passenger
	for _, p := range passengers {
		if p.ID == uint(passengerID) {
			foundPassenger = &p
			break
		}
	}
	if foundPassenger == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Passenger not found"})
	}
	return c.JSON(http.StatusOK, foundPassenger)
}
