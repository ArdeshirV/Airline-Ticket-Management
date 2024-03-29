package persistence

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	db, _ := database.GetDatabaseConnection()
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) GetById(id uint) (*domain.User, error) {
	user := new(domain.User)
	db, _ := database.GetDatabaseConnection()
	tx := db.Debug().Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	db, _ := database.GetDatabaseConnection()
	tx := db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (ur *UserRepository) GeByUsername(username string) (*domain.User, error) {
	user := new(domain.User)
	db, _ := database.GetDatabaseConnection()
	tx := db.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (ur *UserRepository) UpdateById(id uint, newUser *domain.User) (*domain.User, error) {
	db, _ := database.GetDatabaseConnection()
	user := new(domain.User)
	result := db.Debug().First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	user.IsLoginRequired = newUser.IsLoginRequired
	db.Debug().Save(&user)

	return user, nil
}

func (ur *UserRepository) Get(id int) (*domain.User, error) {
	var user domain.User
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&user)

	checkUserExist := db.Debug().Where("ID = ?", id)

	tx := checkUserExist.First(&user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetAll() (*[]domain.User, error) {
	var users []domain.User
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&users)

	checkUserExist := db.Debug().Find(&users)

	if checkUserExist.RowsAffected <= 0 {
		return &users, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&users)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (ur *UserRepository) Delete(id int) error {
	user, err := ur.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&user)
	deleted := db.Debug().Delete(user)
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
