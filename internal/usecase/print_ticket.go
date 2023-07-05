package usecase

import (
	"fmt"

	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/pdf"
)

func CreateTicketAsPDF(id int, TicketFileName string) error {
	title, data, err := GetTicketData(id)
	if err != nil {
		return err
	}
	// TODO: Use --> func GetAirlineLogoByName(name string) (string, error)
	return pdf.CreatePDF(TicketFileName, title, config.Config.App.ImageLogo, data)
}

func GetTicketData(id int) (title string, contents [][]string, err error) {
	err = nil
	title = "The Go Dragons - Team 3 of Quera Software Engineering Bootcamp"
	tr := persistence.NewTicketRepository()
	//tickets, err := tr.GetAll()
	ticket, err := tr.Get(id)
	if err != nil {
		return "", nil, err
	}
	contents = [][]string{
		{"First Name", ticket.Passenger.FirstName}, {"Flight No", ticket.Flight.FlightNo},
		{"Last Name", ticket.Passenger.LastName}, {"Departure", ticket.Flight.Departure.Name}, {"Price", "10"},
		{"National Code", ticket.Passenger.NationalCode}, {"Destination", ticket.Flight.Destination.Name}, {"Price", "10"},
		{"Gender", fmt.Sprintf("%v", ticket.Passenger.Gender)}, {"Terminal", ticket.Flight.Departure.Terminal}, {"Price", "10"},
		{"Birthday", ticket.Passenger.BirthDate.Format("2006/01/02")}, {"Flight Class", fmt.Sprintf("%v", ticket.Flight.FlightClass)}, {"Price", "10"},
	}
	fmt.Printf("%v", contents) // DEBUG
	return title, contents, err
}
