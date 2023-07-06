package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
)

const (
	ParamTime     = "time"
	ParamCityA    = "city_a"
	ParamCityB    = "city_b"
	ParamFlightNo = "flightno"
)

func DataRoute(e *echo.Echo) {
	e.GET("/data", dataHandler)
}

// Test this end-point with these commands:
//
//	              Get all flights: http://localhost:8080/data
//	Get flights with specified ID: http://localhost:8080/data?flightno=VN931
//
// Get flights from A to b in time: http://localhost:8080/data?city_a=New%20York&city_b=Paris&time=2023-06-14
func dataHandler(ctx echo.Context) error {
	flightNo := ctx.QueryParam(ParamFlightNo)
	if flightNo != "" {
		fliteredFlight, err := mock_api.GetFlightsByFlightNo(flightNo)
		if err != nil {
			return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
		}
		return echoJSON(ctx, http.StatusOK, fliteredFlight)
	} else {
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
			}

			filteredFlights, err := mock_api.GetFlightsFromA2B(timeD, cityA, cityB)
			if err != nil {
				return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
			}
			return echoJSON(ctx, http.StatusOK, filteredFlights)
		}
	}
	flights, err := mock_api.GetFlights()
	if err != nil {
		return echoErrorAsJSON(ctx, http.StatusBadRequest, err)
	}
	return echoJSON(ctx, http.StatusOK, flights)
}
