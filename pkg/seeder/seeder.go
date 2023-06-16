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

	_, err := roleRepository.GetByName("user")

	if err != nil {
		newRole := domain.Role{
			Name: "user",
			Description: "normal user",
		}

		_, err = roleRepository.Create(&newRole)
		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
	}

	_, err = userRepository.GetByEmail("admin@gmail.com")

	if err != nil {
		role, _ := roleRepository.GetByName("user")
		encryptedPassword, _ := bcrypt.GenerateFromPassword(
			[]byte("12345678"),
			bcrypt.DefaultCost,
		)
		hashedPassword := string(encryptedPassword)

		newUser := domain.User{
			Email:    "admin@gmail.com", // TODO
			Username: "admin", // TODO
			Password: hashedPassword,
			Phone:    "09035193426", // TODO
			RoleID: role.ID,
		}

		_, err := userRepository.Create(&newUser)

		if err != nil {
			fmt.Printf("could not run seeder: %v\n", err)
		}
	}
	fmt.Print("Seeder finished successfully")

}