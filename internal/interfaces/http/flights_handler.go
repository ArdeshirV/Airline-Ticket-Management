package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	model "github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
)

const (
	BadRequest = "bad request"
	TimeLayout = "2006-01-02_15:04:05"

	ParamFlightNo = "flightno"

	ParamTime  = "time"
	ParamCityA = "city_a"
	ParamCityB = "city_b"

	ParamAscending           = "asc"
	ParamDescending          = "desc"
	ParamSortByPrice         = "sort_by_price"
	ParamSortByDuration      = "sort_by_duration"
	ParamSortByDepartureTime = "sort_by_departure_time"

	ParamDepartureTime           = "departure_time"
	ParamAirlineName             = "airline_name"
	ParamAircraftType            = "aircraft_type"
	ParamRemainingCapacity       = "remaining_capacity"
	ParamRemainingCapacityExists = "remaining_capacity_exists"
)

func DataRoute(e *echo.Echo) {
	e.GET("/flights", flightsHandler)
}

// Test this end-point with these commands:
// http://localhost:8080/flights
// http://localhost:8080/flights?flightno=VN931
// http://localhost:8080/flights?departure_time=2023-10-14_15:30:00
// http://localhost:8080/flights?airline_name=Air%20France
// http://localhost:8080/flights?aircraft_type=Happy%20Airplane
// http://localhost:8080/flights?remaining_capacity=0
// http://localhost:8080/flights?remaining_capacity_exists
// http://localhost:8080/flights?sort_by_price=asc
// http://localhost:8080/flights?sort_by_price=desc
// http://localhost:8080/flights?sort_by_duration=asc
// http://localhost:8080/flights?sort_by_duration=desc
// http://localhost:8080/flights?sort_by_departure_time=asc
// http://localhost:8080/flights?sort_by_departure_time=desc
// http://localhost:8080/flights?sort_by_price=asc&sort_by_departure_time=desc
// http://localhost:8080/data?city_a=New%20York&city_b=Paris&time=2023-06-14
// http://localhost:8080/flights?aircraft_type=Happy%20Airplane&sort_by_departure_time=asc
func flightsHandler(ctx echo.Context) error {
	flightNo := ctx.QueryParam(ParamFlightNo)
	if flightNo != "" {
		filteredFlight, err := mock_api.GetFlightsByFlightNo(flightNo)
		if err != nil {
			return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
		}
		return echoJSON(ctx, http.StatusOK, filteredFlight)
	}

	var err error
	var filteredFlights []model.Flight
	cityA := ctx.QueryParam(ParamCityA)
	cityB := ctx.QueryParam(ParamCityB)
	timeD := ctx.QueryParam(ParamTime)
	if timeD != "" || cityA != "" || cityB != "" {
		errMsg := ""
		dataIsNotEnough := false

		if timeD == "" {
			dataIsNotEnough = true
			errMsg += "'time' is not defined correctly. "
		}
		if cityA == "" {
			dataIsNotEnough = true
			errMsg += "'city_a' is not defined correctly. "
		}
		if cityB == "" {
			dataIsNotEnough = true
			errMsg += "'city_b' is not defined correctly. "
		}
		if dataIsNotEnough {
			return echoStringAsJSON(ctx, http.StatusBadRequest, errMsg)
		} else {
			filteredFlights, err = mock_api.GetFlightsFromA2B(timeD, cityA, cityB)
			if err != nil {
				return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
			}
		}
	} else {
		filteredFlights, err = mock_api.GetFlights()
		if err != nil {
			return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
		}
	}

	departureTime := ctx.QueryParam(ParamDepartureTime)
	if departureTime != "" {
		dateTime, err := time.Parse(TimeLayout, departureTime)
		if err != nil {
			return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
		}
		filteredFlights = filterFlightsByDepartureTime(filteredFlights, dateTime)
	}

	airlineName := ctx.QueryParam(ParamAirlineName)
	if airlineName != "" {
		filteredFlights = filterFlightsByAirlineName(filteredFlights, airlineName)
	}

	aircraftType := ctx.QueryParam(ParamAircraftType)
	if aircraftType != "" {
		filteredFlights = filterFlightsByAircraftType(filteredFlights, aircraftType)
	}

	remainingCapacity := ctx.QueryParam(ParamRemainingCapacity)
	if remainingCapacity != "" {
		remainingCapacityNum, err := strconv.Atoi(remainingCapacity)
		if err != nil {
			return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
		}
		filteredFlights = filterFlightsByRemainingCapacity(filteredFlights, remainingCapacityNum)
	}

	remainingCapacityExists := ctx.QueryParam(ParamRemainingCapacityExists)
	if remainingCapacityExists != "" {
		input := strings.ToLower(strings.TrimSpace(remainingCapacityExists))
		if input == "true" {
			filteredFlights = filterFlightsByRemainingCapacityExists(filteredFlights, true)
		} else if input == "false" {
			filteredFlights = filterFlightsByRemainingCapacityExists(filteredFlights, false)
		} else {
			return echoStringAsJSON(ctx, http.StatusBadRequest, BadRequest)
		}
	}

	sortByPrice := ctx.QueryParam(ParamSortByPrice)
	if sortByPrice != "" {
		if sortByPrice == ParamAscending {
			filteredFlights = sortFlightsByPrice(filteredFlights, true)
		} else if sortByPrice == ParamDescending {
			filteredFlights = sortFlightsByPrice(filteredFlights, false)
		} else {
			return echoStringAsJSON(ctx, http.StatusBadRequest, BadRequest)
		}
	}

	sortByDuration := ctx.QueryParam(ParamSortByDuration)
	if sortByDuration != "" {
		if sortByDuration == ParamAscending {
			filteredFlights = sortFlightsByDuration(filteredFlights, true)
		} else if sortByDuration == ParamDescending {
			filteredFlights = sortFlightsByDuration(filteredFlights, false)
		} else {
			return echoStringAsJSON(ctx, http.StatusBadRequest, BadRequest)
		}
	}

	sortByDepartureTime := ctx.QueryParam(ParamSortByDepartureTime)
	if sortByDepartureTime != "" {
		if sortByDepartureTime == ParamAscending {
			filteredFlights = sortFlightsByDepartureTime(filteredFlights, true)
		} else if sortByDepartureTime == ParamDescending {
			filteredFlights = sortFlightsByDepartureTime(filteredFlights, false)
		} else {
			return echoStringAsJSON(ctx, http.StatusBadRequest, BadRequest)
		}
	}

	return echoJSON(ctx, http.StatusOK, filteredFlights)
}

func filterFlightsByRemainingCapacityExists(flights []model.Flight, cond bool) []model.Flight {
	var filteredFlights []model.Flight
	for _, flight := range flights {
		if cond {
			if flight.RemainingCapacity != 0 {
				filteredFlights = append(filteredFlights, flight)
			}
		} else {
			if flight.RemainingCapacity == 0 {
				filteredFlights = append(filteredFlights, flight)
			}
		}
	}
	return filteredFlights
}

func filterFlightsByDepartureTime(flights []model.Flight, t time.Time) []model.Flight {
	var filteredFlights []model.Flight
	for _, flight := range flights {
		if flight.DepartureTime.Equal(t) {
			filteredFlights = append(filteredFlights, flight)
		}
	}
	return filteredFlights
}

func filterFlightsByAirlineName(flights []model.Flight, airlineName string) []model.Flight {
	var filteredFlights []model.Flight
	for _, flight := range flights {
		// TODO: Check if Airplane and Airline are nil or not
		// It is not necessary when we work with mock-API
		if flight.Airplane.Airline.Name == airlineName {
			filteredFlights = append(filteredFlights, flight)
		}
	}
	return filteredFlights
}

func filterFlightsByAircraftType(flights []model.Flight, aircraftType string) []model.Flight {
	var filteredFlights []model.Flight
	for _, flight := range flights {
		if flight.Airplane.Name == aircraftType {
			filteredFlights = append(filteredFlights, flight)
		}
	}
	return filteredFlights
}

func filterFlightsByRemainingCapacity(flights []model.Flight, remainingCapacity int) []model.Flight {
	var filteredFlights []model.Flight
	for _, flight := range flights {
		if flight.RemainingCapacity == remainingCapacity {
			filteredFlights = append(filteredFlights, flight)
		}
	}
	return filteredFlights
}

func sortFlightsByDuration(flights []model.Flight, asc bool) []model.Flight {
	if asc {
		for i := 0; i < len(flights)-1; i++ {
			for j := i + 1; j < len(flights); j++ {
				if flights[i].ArrivalTime.Sub(flights[i].DepartureTime) > flights[j].ArrivalTime.Sub(flights[j].DepartureTime) {
					flights[i], flights[j] = flights[j], flights[i]
				}
			}
		}
	} else {
		for i := 0; i < len(flights)-1; i++ {
			for j := i + 1; j < len(flights); j++ {
				if flights[i].ArrivalTime.Sub(flights[i].DepartureTime) < flights[j].ArrivalTime.Sub(flights[j].DepartureTime) {
					flights[i], flights[j] = flights[j], flights[i]
				}
			}
		}
	}
	return flights
}

func sortFlightsByPrice(flights []model.Flight, asc bool) []model.Flight {
	if asc {
		for i := 0; i < len(flights)-1; i++ {
			for j := i + 1; j < len(flights); j++ {
				if flights[i].Price > flights[j].Price {
					flights[i], flights[j] = flights[j], flights[i]
				}
			}
		}
	} else {
		for i := 0; i < len(flights)-1; i++ {
			for j := i + 1; j < len(flights); j++ {
				if flights[i].Price < flights[j].Price {
					flights[i], flights[j] = flights[j], flights[i]
				}
			}
		}
	}
	return flights
}

func sortFlightsByDepartureTime(flights []model.Flight, asc bool) []model.Flight {
	if asc {
		for i := 0; i < len(flights)-1; i++ {
			for j := i + 1; j < len(flights); j++ {
				if flights[i].DepartureTime.After(flights[j].DepartureTime) {
					flights[i], flights[j] = flights[j], flights[i]
				}
			}
		}
	} else {
		for i := 0; i < len(flights)-1; i++ {
			for j := i + 1; j < len(flights); j++ {
				if flights[i].DepartureTime.Before(flights[j].DepartureTime) {
					flights[i], flights[j] = flights[j], flights[i]
				}
			}
		}
	}
	return flights
}
