package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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
			found_index := -1
			for index := range mockFlights {
				if mockFlights[index].ID == fs[0].ID {
					found_index = index
					break
				}
			}
			if found_index >= 0 {
				errMsg := fmt.Sprintf("flight not found in mock with this id:%v", fs[0].ID)
				assert.Error(t, errors.New(errMsg))
			} else {
				assert.Equal(t, mockFlights[found_index].FlightNo, fs[0].FlightNo)
			}
		}
	}
	const flightno = "VN931"
	req, err = http.NewRequest(http.MethodGet, "/flights?flightno="+flightno, nil)
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
		if len(fs) > 0 {
			assert.Equal(t, fs[0].FlightNo, flightno)
			found_index := -1
			for index := range mockFlights {
				if mockFlights[index].FlightNo == flightno {
					found_index = index
					break
				}
			}
			if found_index < 0 {
				t.Errorf("flight with flightno:%v not found in mock-api", flightno)
			}
		}
	}
	const (
		destCity    = "Paris"
		sourceCity  = "New York"
		apiTemplate = "/flights?city_a=%s&city_b=%s&time=2023-06-14"
	)
	api := strings.Replace(fmt.Sprintf(apiTemplate, sourceCity, destCity), " ", "%20", -1)
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
		if len(fs) > 0 {
			flight := fs[0]
			assert.Equal(t, sourceCity, flight.Departure.City.Name)
			assert.Equal(t, destCity, flight.Destination.City.Name)
		} else {
			t.Errorf("flight with flightno:%v not found in mock", flightno)
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
