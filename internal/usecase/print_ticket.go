package usecase

import (
	"fmt"

	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/mock_api"
	"github.com/the-go-dragons/final-project/pkg/pdf"
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

func GetTicketData(id int) (title string, logo string, contents [][]string, err error) {
	tr := persistence.NewTicketRepository()
	//tickets, err := tr.GetAll()
	ticket, err := tr.Get(id)
	if err != nil {
		return "", "", nil, err
	}
	logo = ticket.Flight.Airplane.Airline.Logo
	title = fmt.Sprintf("Airline %v, Flights No %v, %v %v, NC %v",
		ticket.Flight.Airplane.Airline.Name,
		ticket.Flight.FlightNo,
		ticket.Passenger.FirstName,
		ticket.Passenger.LastName,
		ticket.Passenger.NationalCode)
	contents = [][]string{
		{"First Name", ticket.Passenger.FirstName}, {"Flight No", ticket.Flight.FlightNo},
		{"Last Name", ticket.Passenger.LastName}, {"Departure", ticket.Flight.Departure.Name},
		{"National Code", ticket.Passenger.NationalCode}, {"Destination", ticket.Flight.Destination.Name},
		{"Gender", fmt.Sprintf("%v", ticket.Passenger.Gender)}, {"Flight Class", fmt.Sprintf("%v", ticket.Flight.FlightClass)},
		{"Birthday", ticket.Passenger.BirthDate.Format("2006/01/02")}, {"Terminal", ticket.Flight.Departure.Terminal},
		{"Seat No" /*ticket.X*/, "7"}, {"Departure Time", ticket.Flight.DepartureTime.Format("2006/01/02 15:04")},
		{"Arrival Time", ticket.Flight.ArrivalTime.Format("2006/01/02 15:04")}, {"Price", fmt.Sprintf("%v", ticket.Flight.Price)},
		{"Airline", ticket.Flight.Airplane.Airline.Name}, {"Airline Logo", ticket.Flight.Airplane.Airline.Logo},
	}
	return title, logo, contents, err
}
