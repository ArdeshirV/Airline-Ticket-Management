package persistance

import (
	"errors"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
	"net/http"
	"strconv"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) New() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) Create(user *domain.User) (*domain.User, error) {
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&user)
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	} else {
		result.Commit()
	}
	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	db, _ := database.GetDatabaseConnection()
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (ur *UserRepository) GeByUsername(username string) (*domain.User, error) {
	user := new(domain.User)
	db, _ := database.GetDatabaseConnection()
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (ur *UserRepository) Update(user *domain.User) (*domain.User, error) {
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&user)

	checkUserExist := db.Debug().Where(&user, "ID = ?", user.ID)
	if checkUserExist.RowsAffected <= 0 {
		return user, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkUserExist.Update("ID", user.ID).Update("Username", user.Username).Update("Password", user.Password)
	tx = tx.Update("Email", user.Email).Update("Phone", user.Phone).Update("CreatedAt", user.CreatedAt)
	tx = tx.Update("Role", user.Role).Update("Passengers", user.Passengers)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedUser := tx.Commit()
		if updatedUser.RowsAffected < 1 {
			return user, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return user, nil
}

func (ur *UserRepository) Get(id int) (*domain.User, error) {
	var user domain.User
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&user)

	checkUserExist := db.Debug().Where(&user, "ID = ?", id)

	if checkUserExist.RowsAffected <= 0 {
		return &user, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkUserExist.Find(&user)

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
	deleted := db.Debug().Delete(user).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
