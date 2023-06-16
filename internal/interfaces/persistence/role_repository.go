package persistence

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type RoleRepository struct{}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (rr *RoleRepository) Create(input *domain.Role) (*domain.Role, error) {
	var role domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&role)

	checkRoleExist := db.Debug().First(&role, "ID = ?", input.ID)

	if checkRoleExist.RowsAffected > 0 {
		return &role, errors.New(strconv.Itoa(http.StatusConflict))
	}

	role.Name = input.Name
	role.Description = input.Description

	addNewRole := db.Debug().Create(&role).Commit()

	if addNewRole.RowsAffected < 1 {
		return &role, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &role, nil
}

func (rr *RoleRepository) GetById(id int) (*domain.Role, error) {
	var role domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&role)

	checkRoleExist := db.Debug().Where(&role, "ID = ?", id)

	if checkRoleExist.RowsAffected <= 0 {
		return &role, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkRoleExist.Find(&role)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (rr *RoleRepository) GetByName(name string) (*domain.Role, error) {
	role := new(domain.Role)
	db, _ := database.GetDatabaseConnection()
	db.Where("name = ?", name).First(&role)
	if role.ID == 0 {
		return nil, errors.New("Role not found")
	}
	return role, nil
}

func (a *RoleRepository) Update(input *domain.Role) (*domain.Role, error) {
	var role domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&role)

	checkRoleExist := db.Debug().Where(&role, "ID = ?", input.ID)

	if checkRoleExist.RowsAffected <= 0 {
		return &role, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkRoleExist.Update("ID", input.ID).Update("Name", input.Name).Update("Description", input.Description)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedRole := tx.Commit()
		if updatedRole.RowsAffected < 1 {
			return &role, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &role, nil
}

// func (a *RoleRepository) Get(id int) (*domain.Role, error) {
// 	var role domain.Role
// 	db, _ := database.GetDatabaseConnection()
// 	db = db.Model(&role)

// 	checkRoleExist := db.Debug().Where(&role, "ID = ?", id)

// 	if checkRoleExist.RowsAffected <= 0 {
// 		return &role, errors.New(strconv.Itoa(http.StatusNotFound))
// 	}

// 	tx := checkRoleExist.Find(&role)

// 	if err := tx.Error; err != nil {
// 		return nil, err
// 	}

// 	return &role, nil
// }

func (a *RoleRepository) GetAll() (*[]domain.Role, error) {
	var roles []domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&roles)

	checkRoleExist := db.Debug().Find(&roles)

	if checkRoleExist.RowsAffected <= 0 {
		return &roles, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&roles)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &roles, nil
}

func (a *RoleRepository) Delete(id int) error {
	role, err := a.GetById(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&role)
	deleted := db.Debug().Delete(role).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
