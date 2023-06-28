package http

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	models "github.com/the-go-dragons/final-project/internal/domain"
	_ "github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
)

func TestFlightsHandler(t *testing.T) {
	e := echo.New()
	mockFlights, err := mock_api.GetFlights()
	if err != nil {
		t.Errorf("mock_api.GetFlights() returns: %v", err)
	}
	req, err := http.NewRequest(http.MethodGet, "/flights", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, flightsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fs, err := ConvertJSON2Flights(rec.Body.String())
		if err != nil {
			t.Errorf("Error failed to convert rec.Body from JSON to []Flight: %v", err)
		}
		assert.Equal(t, len(mockFlights), len(fs))
		if len(fs) > 0 {
			assert.Equal(t, mockFlights[0].ID, fs[0].ID)
			assert.Equal(t, mockFlights[0].FlightNo, fs[0].FlightNo)
		}
	}
	req, err = http.NewRequest(http.MethodGet, "/flights?flightno=LA882", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(t, flightsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fs, err := ConvertJSON2Flights(rec.Body.String())
		if err != nil {
			t.Errorf("Error failed to convert rec.Body from JSON to []Flight: %v", err)
		}
		assert.Equal(t, len(mockFlights), len(fs))
		if len(fs) > 0 {
			assert.Equal(t, mockFlights[0].ID, fs[0].ID)
			assert.Equal(t, mockFlights[0].FlightNo, fs[0].FlightNo)
		}
	}
	api := "/flights?city_a=New%20York&city_b=Paris&time=2023-06-14"
	req, err = http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(t, flightsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fs, err := ConvertJSON2Flights(rec.Body.String())
		if err != nil {
			t.Errorf("Error failed to convert rec.Body from JSON to []Flight: %v", err)
		}
		assert.Equal(t, len(mockFlights), len(fs))
		if len(fs) > 0 {
			assert.Equal(t, mockFlights[0].ID, fs[0].ID)
			assert.Equal(t, mockFlights[0].FlightNo, fs[0].FlightNo)
		}
	}
}

func ConvertJSON2Flights(JSONData string) ([]models.Flight, error) {
	var flights []models.Flight
	err := json.Unmarshal([]byte(JSONData), &flights)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal flight data: %v", err)
	}
	return flights, nil
}
