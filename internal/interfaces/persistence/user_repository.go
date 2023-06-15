package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	db := database.DBConn
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	db := database.DBConn
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (ur *UserRepository) GeByUsername(username string) (*domain.User, error) {
	user := new(domain.User)
	db := database.DBConn
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, errors.New("User not found")
	}
	return user, nil
}
