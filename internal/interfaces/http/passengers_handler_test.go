package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "github.com/the-go-dragons/final-project/internal/domain"
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
	/*assert.Equal(t, http.StatusOK, rec.Code)

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
	assert.Equal(t, "123 Example Street", passengers[0].Address)*/
}
