package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
)

type APIResponse struct {
	Message string `json: "message"`
}

const (
	ParamTime = "time"
	ParamCityA = "city_a"
	ParamCityB = "city_b"
	ParamCommand = "command"
	ParamFlightNo = "flightno"
)

func FlightsRoute(e *echo.Echo) {
	e.GET("/flights", flightsHandler)
}

func flightsHandler(ctx echo.Context) error {
	flightNo := ctx.QueryParam(ParamFlightNo)
	if flightNo != "" {
		fliteredFlight, err := mock_api.GetFlightsByFlightNo(flightNo)
		if err != nil {
			return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
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
			} else {
				filteredFlights, err := mock_api.GetFlightsFromA2B(timeD, cityA, cityB)
				if err != nil {
					return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
				}
				return echoJSON(ctx, http.StatusOK, filteredFlights)
			}
		} else {
			data, err := mock_api.GetFlights()
			if err != nil {
				return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: fmt.Sprintf("%v", err) })
			}
			return echoJSON(ctx, http.StatusOK, data)
		}
	}
	return echoJSON(ctx, http.StatusBadRequest, APIResponse{ Message: "Bad request" })
}

func echoJSON(ctx echo.Context, status int, data interface{}) error {
	// TODO: Add config.IsDebugModeEnabled() to config.go and then uncomment codes
	//if config.IsDebugModeEnabled() {
		return ctx.JSONPretty(status, data, "    ")
	//} else {
		//return ctx.JSON(status, data)
	//}
}
