package usecase

import (
	"fmt"
	"strings"

	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
	"github.com/the-go-dragons/final-project/pkg/pdf"
)

const (
	paid           = "paid"
	birthdayLayout = "2006/01/02"
	dateTimeLayout = "2006/01/02 15:04"
)

func CreateTicketAsPDF(id int, TicketFileName string) error {
	title, logoName, data, err := GetTicketData(id)
	if err != nil {
		return err
	}
	logoFile, err := mock_api.GetAirlineLogoByName(logoName)
	if err != nil {
		logoFile = config.Config.App.ImageLogo
	}
	return pdf.CreatePDF(TicketFileName, title, logoFile, data)
}

func GetTicketData(id int) (string, string, [][]string, error) {
	ticketRepo := persistence.NewTicketRepository()
	ticket, err := ticketRepo.Get(id)
	if err != nil {
		return "", "", nil, err
	}
	if strings.ToLower(strings.TrimSpace(ticket.PaymentStatus)) != paid {
		err = fmt.Errorf("the ticket price with ID:%v has not been paid yet", ticket.ID)
		return "", "", nil, err
	}
	flightRepo := persistence.NewFlightRepository()
	flight, err := flightRepo.Get(int(ticket.FlightID))
	if err != nil {
		return "", "", nil, err
	}
	passengerRepo := persistence.NewPassengerRepository()
	passenger, err := passengerRepo.Get(int(ticket.PassengerID))
	if err != nil {
		return "", "", nil, err
	}
	airplaneRepo := persistence.NewAirplaneRepository()
	airplane, err := airplaneRepo.Get(int(flight.AirplaneID))
	if err != nil {
		return "", "", nil, err
	}
	airlineRepo := persistence.NewAirlineRepsoitory()
	airline, err := airlineRepo.Get(int(airplane.AirlineID))
	if err != nil {
		return "", "", nil, err
	}
	airportRepo := persistence.NewAirportRepository()
	departure, err := airportRepo.Get(int(flight.DepartureID))
	if err != nil {
		return "", "", nil, err
	}
	destination, err := airportRepo.Get(int(flight.DestinationID))
	if err != nil {
		return "", "", nil, err
	}
	cityRepo := persistence.NewCityRepository()
	departureCity, err := cityRepo.Get(int(departure.CityID))
	if err != nil {
		return "", "", nil, err
	}
	destinationCity, err := cityRepo.Get(int(destination.CityID))
	if err != nil {
		return "", "", nil, err
	}
	title := fmt.Sprintf("Airline: %v, Flights No: %v, %v %v, NC %v",
		airline.Name,
		flight.FlightNo,
		passenger.FirstName,
		passenger.LastName,
		passenger.NationalCode)
	contents := [][]string{
		{"First Name", passenger.FirstName},
		{"Last Name", passenger.LastName},
		{"National Code", passenger.NationalCode},
		//{"Gender", fmt.Sprintf("%v", passenger.Gender)},
		//{"Birthday", passenger.BirthDate.Format(birthdayLayout)},
		//{"Phone", passenger.Phone},
		//{"Email", passenger.Email},
		//{"Address", passenger.Address},
		{"", ""},
		{"Flight No", flight.FlightNo},
		{"Departure/Destination", fmt.Sprintf("%v --> %v", departureCity.Name, destinationCity.Name)},
		{"Class", fmt.Sprintf("%v", flight.FlightClass)},
		{"Departure Time", flight.DepartureTime.Format(dateTimeLayout)},
		//{"Arrival Time", flight.ArrivalTime.Format(dateTimeLayout)},
		//{"Remaining Capacity", fmt.Sprintf("%v", flight.RemainingCapacity)},
		//{"", airline.Name},
		//{"", airline.Logo},
		{"Departure Airport", departure.Name},
		{"Destination Airport", destination.Name},
		//{"Terminal", departure.Terminal},
		//{"Cancel Condition", flight.CancelCondition},
		{"Price", fmt.Sprintf("%v", flight.Price)},
	}
	//ShowValue("title", title)
	//ShowValue("contents", contents)
	return title, airline.Logo, contents, err
}

func ShowValue(title string, value interface{}) {
	if config.IsDebugMode() {
		fmt.Printf("  %s: \033[1;33m%v\033[0;0m\n", title, value)
	}
}
