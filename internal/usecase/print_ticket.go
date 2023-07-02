package usecase

import (
	"fmt"

	"github.com/the-go-dragons/final-project/internal/domain"
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
	tickets, err := tr.GetAll()
	if err != nil {
		return "", nil, err
	}
	var ticket *domain.Ticket = nil
	if len(*tickets) == 0 {
		ticket = createFakeTicket() // TODO: Remove this line when database works
	} else {
		for _, t := range *tickets {
			if uint(t.ID) == uint(id) {
				ticket = &t
			}
		}
	}
	if ticket == nil {
		return "", nil, fmt.Errorf("ticket with id:'%v' not found", id)
	}
	fmt.Printf("\\nticket: %v\\n", ticket)

	contents = [][]string{
		{"Name", "Ardeshir1"}, {"Price", "10"},
		{"Name", "Ardeshir2"}, {"Price", "10"},
		{"Name", "Ardeshir3"}, {"Price", "10"},
		{"Name", "Ardeshir4"}, {"Price", "10"},
		{"Name", "Ardeshir5"}, {"Price", "10"},
		{"Name", "Ardeshir6"}, {"Price", "10"},
		{"Name", "Ardeshir7"}, {"Price", "10"},
		{"Name", "Ardeshir8"}, {"Price", "10"},
		{"Name", "Ardeshir9"}, {"Price", "10"},
		{"Name", "Ardeshir10"}, {"Price", "10"},
	}
	return title, contents, err
}

func createFakeTicket() *domain.Ticket {
	if config.IsDebugMode() {
		fmt.Printf("Debug: Create fake ticket")
	}
	return &domain.Ticket{}
	/*&domain.Ticket{
		FlightID: 1,
		Flight: Flight {
			FlightNo: "",
			DepartureID: 1,
			Departure: Airport {
			},
			DestinationID: 1,
			Destination: Airport {
			},
			DepartureTime     time.Time   `json:"departuretime"`
			ArrivalTime       time.Time   `json:"arrivaltime"`
			AirplaneID        uint        `json:"airplaneid"`
			Airplane          Airplane    `json:"airplane"`
			FlightClass       FlightClass `json:"flightclass"`
			Price             int         `json:"price"`
			RemainingCapacity int         `json:"remainingcapacity"`
			CancelCondition   string      `json:"cancelcondition"`
		},
		PassengerID: 1,
		Passenger: Passenger {
		},
		PaymentID: 1,
		Payment: Payment {
		},
		UserID: 1,
		User: User {
		},
		PaymentStatus: "",
		Refund: true,
	}*/
}
