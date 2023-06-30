package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/the-go-dragons/final-project/internal/domain"
	"gorm.io/gorm"
)

func TestAddPassenger(t *testing.T) {
	e := echo.New()

	requestBody := map[string]interface{}{
		"firstName":    "Akbar",
		"lastName":     "Akbari",
		"nationalCode": "1234567890",
		"email":        "akbar.akabari@example.com",
		"gender":       "male",
		"birthDate":    "1990-01-01",
		"phone":        "1234567890",
		"address":      "123 Example Street",
	}
	requestJSON, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/passengers", bytes.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddPassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"message":"Passenger added successfully","passengerId":1}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	assert.Len(t, passengers, 1)
	assert.Equal(t, uint(1), passengers[0].ID)
	assert.Equal(t, "Akbar", passengers[0].FirstName)
	assert.Equal(t, "Akbari", passengers[0].LastName)
	assert.Equal(t, "1234567890", passengers[0].NationalCode)
	assert.Equal(t, "akbar.akbari@example.com", passengers[0].Email)
	assert.Equal(t, "male", passengers[0].Gender)
	assert.Equal(t, "1990-01-01", passengers[0].BirthDate)
	assert.Equal(t, "1234567890", passengers[0].Phone)
	assert.Equal(t, "123 Example Street", passengers[0].Address)
}

func TestAddPassenger_InvalidPayload(t *testing.T) {
	e := echo.New()

	requestBody := map[string]interface{}{
		"invalidField": "invalidValue",
	}
	requestJSON, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/passengers", bytes.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddPassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	expectedResponse := `{"error":"Invalid request payload"}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	assert.Len(t, passengers, 0)
}

func TestDeletePassenger(t *testing.T) {
	e := echo.New()

	passengerID := 1

	req := httptest.NewRequest(http.MethodDelete, "/passengers/"+strconv.Itoa(passengerID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(passengerID))

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: uint(passengerID)}, FirstName: "Akbar", LastName: "Akbari"},
	}

	err := DeletePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"message":"Passenger deleted successfully"}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	assert.Empty(t, passengers)
}

func TestDeletePassenger_InvalidID(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/passengers/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")

	err := DeletePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	expectedResponse := `{"error":"Invalid passenger ID"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestDeletePassenger_PassengerNotFound(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/passengers/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: 1}, FirstName: "Akbar", LastName: "Akbari"},
	}

	err := DeletePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	expectedResponse := `{"error":"Passenger not found"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestUpdatePassenger(t *testing.T) {
	e := echo.New()

	passengerID := 1
	payload := map[string]interface{}{
		"firstName":    "Akbar",
		"lastName":     "Akbari",
		"nationalCode": "1234567890",
		"email":        "akbar.akbari@example.com",
		"gender":       "male",
		"birthDate":    "1990-01-01",
		"phone":        "1234567890",
		"address":      "123 Enqelab Street",
	}
	payloadBytes, _ := json.Marshal(payload)
	payloadReader := bytes.NewReader(payloadBytes)

	req := httptest.NewRequest(http.MethodPut, "/passengers/"+strconv.Itoa(passengerID), payloadReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(passengerID))

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: uint(passengerID)}, FirstName: "Asqar", LastName: "Asqari"},
	}

	err := UpdatePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"message":"Passenger updated successfully"}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	updatedPassenger := passengers[0]
	assert.Equal(t, payload["firstName"].(string), updatedPassenger.FirstName)
	assert.Equal(t, payload["lastName"].(string), updatedPassenger.LastName)
	assert.Equal(t, payload["nationalCode"].(string), updatedPassenger.NationalCode)
	assert.Equal(t, payload["email"].(string), updatedPassenger.Email)
	assert.Equal(t, payload["gender"].(int), int(updatedPassenger.Gender))
	assert.Equal(t, payload["birthDate"].(string), updatedPassenger.BirthDate.Format("2006-01-02"))
	assert.Equal(t, payload["phone"].(string), updatedPassenger.Phone)
	assert.Equal(t, payload["address"].(string), updatedPassenger.Address)
}

func TestUpdatePassenger_InvalidID(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/passengers/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")

	err := UpdatePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	expectedResponse := `{"error":"Invalid passenger ID"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestUpdatePassenger_PassengerNotFound(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/passengers/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: 1}, FirstName: "Asqar", LastName: "Asqari"},
	}

	err := UpdatePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	expectedResponse := `{"error":"Passenger not found"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestUpdatePassenger_InvalidPayload(t *testing.T) {
	e := echo.New()

	passengerID := 1
	invalidPayload := map[string]interface{}{
		"invalidField": "value",
	}
	invalidPayloadBytes, _ := json.Marshal(invalidPayload)
	invalidPayloadReader := bytes.NewReader(invalidPayloadBytes)

	req := httptest.NewRequest(http.MethodPut, "/passengers/"+strconv.Itoa(passengerID), invalidPayloadReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(passengerID))

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: uint(passengerID)}, FirstName: "Asqar", LastName: "Asqari"},
	}

	err := UpdatePassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	expectedResponse := `{"error":"Invalid request payload"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestGetPassenger(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/passengers/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: 1}, FirstName: "Akbar", LastName: "Akbari"},
		{Model: gorm.Model{ID: 2}, FirstName: "Asqar", LastName: "Asqari"},
	}

	err := GetPassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"ID":1,"FirstName":"Akbar","LastName":"Akbari"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestGetPassenger_InvalidID(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/passengers/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")

	err := GetPassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	expectedResponse := `{"error":"Invalid passenger ID"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}

func TestGetPassenger_PassengerNotFound(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/passengers/3", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	passengers = []domain.Passenger{
		{Model: gorm.Model{ID: 1}, FirstName: "Akbar", LastName: "Akbari"},
		{Model: gorm.Model{ID: 2}, FirstName: "Asqar", LastName: "Asqari"},
	}

	err := GetPassenger(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	expectedResponse := `{"error":"Passenger not found"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}
