package mock_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	models "github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/config"
)

type command string

const (
	airlineLogoFileName = "pdf/airline_logo.png"
	APIGetLogo          = "http://localhost:%d/airline?logo_name=%s"
	APICities           = "http://localhost:%d/cities"
	APIAirplanes        = "http://localhost:%d/airplanes"
	APIDepartureDates   = "http://localhost:%d/departure_dates"
	APIFlights          = "http://localhost:%d/flights"
	APIFlightByFlightNo = "http://localhost:%d/flights?flightno=%s"
	APIFlightsFromA2B   = "http://localhost:%d/flights?city_a=%s&city_b=%s&time=%s"

	CommandReturn    command = "return"
	CommandReserve   command = "reserve"
	APIFlightReserve         = "http://localhost:%d/reserve_flight?flightno=%s&command=%s"
)

type ReserveResponse struct {
	Message           string `json:"message"`
	FlightNo          string `json:"flightno"`
	Capacity          int    `json:"capacity"`
	RemainingCapacity int    `json:"remainingcapacity"`
}

func GetAirlineLogoByName(name string) (string, error) {
	url := fmt.Sprintf(APIGetLogo, config.Config.Mock.Port, string(name))
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	err = os.WriteFile(airlineLogoFileName, data, 0666)
	if err != nil {
		return "", err
	}
	return airlineLogoFileName, nil
}

func SetRemainingCapacity(flightNo string, cmd command) (resp *ReserveResponse, err error) {
	data := resp
	api := fmt.Sprintf(APIFlightReserve, config.Config.Mock.Port, flightNo, string(cmd))
	api = strings.Replace(api, " ", "%20", len(api))
	JSONResponse, err := ReadJSONFromAPIwithPOST(api, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return data, nil
}

func GetFlightsFromA2B(timeD, cityA, cityB string) (flights []models.Flight, err error) {
	data := &flights
	api := fmt.Sprintf(APIFlightsFromA2B, config.Config.Mock.Port, cityA, cityB, timeD)
	api = strings.Replace(api, " ", "%20", len(api))
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetFlightsByFlightNo(flightNo string) (flight []models.Flight, err error) {
	data := &flight
	api := fmt.Sprintf(APIFlightByFlightNo, config.Config.Mock.Port, flightNo)
	api = strings.Replace(api, " ", "%20", len(api))
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetAirplanes() (airplanes []models.Airplane, err error) {
	data := &airplanes
	api := fmt.Sprintf(APIAirplanes, config.Config.Mock.Port)
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetFlights() (flights []models.Flight, err error) {
	data := &flights
	api := fmt.Sprintf(APIFlights, config.Config.Mock.Port)
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetCities() (cities []models.City, err error) {
	data := &cities
	api := fmt.Sprintf(APICities, config.Config.Mock.Port)
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetDepartureDates() (times []time.Time, err error) {
	data := &times
	api := fmt.Sprintf(APIDepartureDates, config.Config.Mock.Port)
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func ReadJSONFromAPIwithGET(api string) (string, error) {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bodyBytes), err
}

func ReadJSONFromAPIwithPOST(api string, data interface{}) (string, error) {
	jsonReq, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(api, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bodyBytes), err
}

func Normalize(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}
