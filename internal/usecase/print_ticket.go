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
	fmt.Printf("\n\033[0mtitle = %v\ndata = %v\033[0m\n\n", title, data)
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
		{"Last Name", ticket.Passenger.LastName}, {"Departure", ticket.Flight.Departure.Name},
		{"National Code", ticket.Passenger.NationalCode}, {"Destination", ticket.Flight.Destination.Name},
		{"Gender", fmt.Sprintf("%v", ticket.Passenger.Gender)}, {"Flight Class", fmt.Sprintf("%v", ticket.Flight.FlightClass)},
		{"Birthday", ticket.Passenger.BirthDate.Format("2006/01/02")}, {"Terminal", ticket.Flight.Departure.Terminal},
	}
	fmt.Printf("\n\n\033[0;36m%v\033[0m\n\n", contents) // DEBUG
	return title, contents, err
}
