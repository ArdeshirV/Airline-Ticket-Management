package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type UserRepo struct {
	db *gorm.DB
}

func (a *UserRepo) New(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (a *UserRepo) Save(input *domain.User) (*domain.User, error) {
	var user domain.User
	db := a.db.Model(&user)

	checkUserExist := db.Debug().First(&user, "ID = ?", input.ID)

	if checkUserExist.RowsAffected > 0 {
		return &user, errors.New(strconv.Itoa(http.StatusConflict))
	}

	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email
	user.Phone = input.Phone
	user.CreatedAt = input.CreatedAt
	user.Role = input.Role
	user.Passengers = input.Passengers

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		return &user, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &user, nil
}

func (a *UserRepo) Update(input *domain.User) (*domain.User, error) {
	var user domain.User
	db := a.db.Model(&user)

	checkUserExist := db.Debug().Where(&user, "ID = ?", input.ID)

	if checkUserExist.RowsAffected <= 0 {
		return &user, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkUserExist.Update("ID", input.ID).Update("Username", input.Username).Update("Password", input.Password)
	tx = tx.Update("Email", input.Email).Update("Phone", input.Phone).Update("CreatedAt", input.CreatedAt)
	tx = tx.Update("Role", input.Role).Update("Passengers", input.Passengers)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedUser := tx.Commit()
		if updatedUser.RowsAffected < 1 {
			return &user, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &user, nil
}

func (a *UserRepo) Get(id int) (*domain.User, error) {
	var user domain.User
	db := a.db.Model(&user)

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

func (a *UserRepo) GetAll() (*[]domain.User, error) {
	var users []domain.User
	db := a.db.Model(&users)

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

func (a *UserRepo) Delete(id int) error {
	user, err := a.Get(id)
	if err != nil {
		return err
	}
	db := a.db.Model(&user)
	deleted := db.Debug().Delete(user).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
