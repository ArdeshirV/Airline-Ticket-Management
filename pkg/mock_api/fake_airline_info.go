package mock_api

import (
	"io"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	"net/http"

	models "github.com/the-go-dragons/final-project/internal/domain"
)

type command string

const (
	APICities           = "http://localhost:3000/cities"
	APIAirplanes        = "http://localhost:3000/airplanes"
	APIDepartureDates   = "http://localhost:3000/departure_dates"
	APIFlights          = "http://localhost:3000/flights"
	APIFlightByFlightNo = "http://localhost:3000/flights?flightno=%s"
	APIFlightsFromA2B   = "http://localhost:3000/flights?city_a=%s&city_b=%s&time=%s"

	CommandReturn    command = "return"
	CommandReserve   command = "reserve"
	APIFlightReserve         = "http://localhost:3000/reserve_flight?flightno=%s&command=%s"
)

type ReserveResponse struct {
	Message           string `json:"message"`
	FlightNo          string `json:"flightno"`
	Capacity          int    `json:"capacity"`
	RemainingCapacity int    `json:"remainingcapacity"`
}

func SetRemainingCapacity(flightNo string, cmd command) (resp *ReserveResponse, err error) {
	data := resp
	api := fmt.Sprintf(APIFlightReserve, flightNo, string(cmd))
	api = strings.Replace(api, " ", "%20", len(api))
	JSONResponse, err := ReadJSONFromAPIwithPOST(api, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data, nil
}

func GetFlightsFromA2B(timeD, cityA, cityB string) (flights []models.Flight, err error) {
	data := &flights
	api := fmt.Sprintf(APIFlightsFromA2B, cityA, cityB, timeD)
	api = strings.Replace(api, " ", "%20", len(api))
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetFlightsByFlightNo(flightNo string) (flight []models.Flight, err error) {
	data := &flight
	api := fmt.Sprintf(APIFlightByFlightNo, flightNo)
	api = strings.Replace(api, " ", "%20", len(api))
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetAirplanes() (airplanes []models.Airplane, err error) {
	data := &airplanes
	api := APIAirplanes
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetFlights() (flights []models.Flight, err error) {
	data := &flights
	api := APIFlights
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetCities() (cities []models.City, err error) {
	data := &cities
	api := APICities
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data, nil
}

func GetDepartureDates() (times []time.Time, err error) {
	data := &times
	api := APIDepartureDates
	JSONResponse, err := ReadJSONFromAPIwithGET(api)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(JSONResponse), &data)
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

func AreDatesEqual(a, b time.Time) bool {
	return a.Day() == b.Day() && a.Month() == b.Month() && a.Year() == b.Year()
}
