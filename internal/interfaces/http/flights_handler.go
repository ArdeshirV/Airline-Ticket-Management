package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
)

type Flights []domain.Flight

type APIResponse struct {
	Message string `json: "message"`
}

const (
	ParamTime = "time"
	ParamCityA = "city_a"
	ParamCityB = "city_b"
	ParamCommand = "command"
	ParamFlightNo = "flightno"
	ParamMinimumCapacity = "min_capacity"
	ParamDepartureDateTime = "depature_datetime"
	ParamArriveDateTime = "arrive_datetime"
	ParamAirplane = "airplane"
	ParamAirline = "airline"
)

func FlightsRoute(e *echo.Echo) {
	e.GET("/flights", flightsHandler)
}

func flightsHandler(ctx echo.Context) error {
	data := make(Flights, 0)

	flightNo := ctx.QueryParam(ParamFlightNo)
	minCapacity := ctx.QueryParam(ParamMinimumCapacity)
	depatureDatetime := ctx.QueryParam(ParamDepartureDateTime)
	arriveDatetime := ctx.QueryParam(ParamArriveDateTime)
	airplane := ctx.QueryParam(ParamAirplane)
	airline := ctx.QueryParam(ParamAirline)

	if flightNo != "" {
		fliteredFlight, err := mock_api.GetFlightsByFlightNo(flightNo)
		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}
		data = fliteredFlight
	}

	cityA := ctx.QueryParam(ParamCityA)
	cityB := ctx.QueryParam(ParamCityB)
	timeD := ctx.QueryParam(ParamTime)

	if timeD != "" || cityA != "" || cityB != "" {
		errMsg := ""
		dataIsNotEnough := false

		if timeD == "" {
			dataIsNotEnough = true
			errMsg += "\"time\" is not defined correctly. "
		}

		if cityA == "" {
			dataIsNotEnough = true
			errMsg += "\"city_a\" is not defined correctly. "
		}

		if cityB == "" {
			dataIsNotEnough = true
			errMsg += "\"city_b\" is not defined correctly. "
		}

		if dataIsNotEnough {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: errMsg })
		}

		filteredFlights, err := mock_api.GetFlightsFromA2B(timeD, cityA, cityB)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}

		data = filteredFlights
	}

	allFlights, err := mock_api.GetFlights()

	data = allFlights

	if minCapacity != "" {
			
		numberMinCapacity, err := strconv.Atoi(minCapacity)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}

		data, err = data.FilterFlightsByMinimumCapacity(numberMinCapacity)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}
	}

	if airplane != "" {
			
		airplaneId, err := strconv.Atoi(airplane)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}

		data, err = data.FilterFlightsByAirlineId(airplaneId)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}
	}

	if airline != "" {
			
		airlineId, err := strconv.Atoi(airline)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}

		data, err = data.FilterFlightsByAirlineId(airlineId)

		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
		}
	}

	if depatureDatetime != "" || arriveDatetime != "" {
		data, err = data.FilterFlightsByDepatureTimeAndArriveTime(depatureDatetime, arriveDatetime)
	}

	if err != nil {
		return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
	}

	return echoJSON(ctx, http.StatusOK, data)
}

func echoJSON(ctx echo.Context, status int, data interface{}) error {
	// TODO: Add config.IsDebugModeEnabled() to config.go and then uncomment codes
	//if config.IsDebugModeEnabled() {
		return ctx.JSONPretty(status, data, "    ")
	//} else {
		//return ctx.JSON(status, data)
	//}
}

func (flights Flights) FilterFlightsByMinimumCapacity(minimumCapacity int) (Flights, error) {

	filteredFlights := make(Flights, 0)

	for _, flight := range flights {
		if flight.RemainingCapacity >= minimumCapacity {
			filteredFlights = append(filteredFlights, flight)
		}
	}

	return filteredFlights, nil
}

func (flights Flights) FilterFlightsByDepatureTimeAndArriveTime(depatureDatetime string, arriveDateTime string) (Flights, error) {
	filteredFlights := make(Flights, 0)

	var parsedDepatureDatetime time.Time
	var parsedArriveDatetime time.Time
	var err error

	if depatureDatetime != "" {
		parsedDepatureDatetime, err = time.Parse("2006-01-02T15:04:05Z", depatureDatetime)
		if err != nil {
			return nil, err
		}
	}

	if arriveDateTime != "" {
		parsedArriveDatetime, err = time.Parse("2006-01-02T15:04:05Z", arriveDateTime)
		if err != nil {
			return nil, err
		}
	}

	for _, flight := range flights {
		if (depatureDatetime == "" || flight.DepartureTime == parsedDepatureDatetime) &&
			(arriveDateTime == "" || flight.ArrivalTime == parsedArriveDatetime) {
			filteredFlights = append(filteredFlights, flight)
		}
	}

	return filteredFlights, nil
}

func (flights Flights) FilterFlightsByAirplaneId(airplaneId int) (Flights, error) {

	filteredFlights := make(Flights, 0)

	for _, flight := range flights {
		if int(flight.AirplaneID) == airplaneId {
			filteredFlights = append(filteredFlights, flight)
		}
	}

	return filteredFlights, nil
}


func (flights Flights) FilterFlightsByAirlineId(airlineId int) (Flights, error) {

	filteredFlights := make(Flights, 0)

	for _, flight := range flights {
		if int(flight.Airplane.AirlineID) == airlineId {
			filteredFlights = append(filteredFlights, flight)
		}
	}

	return filteredFlights, nil
}
