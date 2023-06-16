package usecase

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repository *persistence.UserRepository
}

func NewUserUsecase(repository *persistence.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
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

	return ur.repository.Create(user)
}

func (ur *UserUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return ur.repository.GetByEmail(email)
}

func (ur *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return ur.repository.GeByUsername(username)
}
