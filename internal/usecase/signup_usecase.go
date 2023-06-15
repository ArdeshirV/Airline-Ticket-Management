package usecase

import (
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
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
	return ur.repository.Create(user)
}

func (ur *UserUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return ur.repository.GetByEmail(email)
}

func (ur *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return ur.repository.GeByUsername(username)
}
