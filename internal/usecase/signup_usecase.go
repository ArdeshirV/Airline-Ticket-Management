package usecase

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository *persistence.UserRepository
	roleRepository *persistence.RoleRepository
}

func NewUserUsecase(repository *persistence.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repository,
	}
}

func (ur *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	// Hash the password
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, errors.New("Cant hash password")
	}
	user.Password = string(encryptedPassword)

	return ur.userRepository.Create(user)
}

func (ur *UserUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return ur.userRepository.GetByEmail(email)
}

func (ur *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return ur.userRepository.GeByUsername(username)
}
