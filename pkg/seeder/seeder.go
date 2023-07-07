package seeder

import (
	"fmt"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"golang.org/x/crypto/bcrypt"
)

func Run() {
	fmt.Print("Seeder runner started")
	userRepository := persistence.NewUserRepository()
	roleRepository := persistence.NewRoleRepository()
	ticketRepository := persistence.NewTicketRepository()
	flightRepository := persistence.NewFlightRepository()
	airplaneRepo := persistence.NewAirplaneRepository()
	airlineRepo := persistence.NewAirlineRepsoitory()

	_, err := roleRepository.GetByName("user")

	if err != nil {
		newRole := domain.Role{
			Name:        "user",
			Description: "normal user",
		}

		_, err = roleRepository.Create(&newRole)
		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
	}

	_, err = roleRepository.GetByName("admin")
	if err != nil {

		newRole := domain.Role{
			Name:        "admin",
			Description: "admin user",
		}

		_, err = roleRepository.Create(&newRole)
		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
	}

	user, err := userRepository.GetByEmail("admin@gmail.com")

	if err != nil {
		role, _ := roleRepository.GetByName("user")
		encryptedPassword, _ := bcrypt.GenerateFromPassword(
			[]byte("12345678"), // TODO should read from db
			bcrypt.DefaultCost,
		)
		hashedPassword := string(encryptedPassword)

		passengers := make([]domain.Passenger, 0)

		newUser := domain.User{
			Email:      "admin@gmail.com", // TODO should read from db
			Username:   "admin",           // TODO should read from db
			Password:   hashedPassword,
			Phone:      "09035193426", // TODO should read from db
			RoleID:     role.ID,
			Passengers: passengers,
		}

		user, err = userRepository.Create(&newUser)

		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
	}

	_, err = ticketRepository.GetByUserId(uint(user.ID))
	if err != nil {
		airline := &domain.Airline{Name: "Test"}
		airlineRepo.Create(airline)
		airplane := &domain.Airplane{AirlineID: airline.ID}
		airplane, err = airplaneRepo.Create(airplane)
		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
		println("ID>>>>>", airplane.ID)
		newflight, err := flightRepository.Create(&domain.Flight{
			FlightNo:          "10",
			RemainingCapacity: 10,
			AirplaneID:        airplane.ID,
		})
		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}

		newTicket := domain.Ticket{
			User:     *user,
			UserID:   uint(user.ID),
			Flight:   *newflight,
			FlightID: newflight.ID,
		}

		_, err = ticketRepository.Create(&newTicket)
		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
	}

	fmt.Print("Seeder finished successfully")

}
