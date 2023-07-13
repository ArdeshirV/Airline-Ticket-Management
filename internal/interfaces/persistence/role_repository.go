package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type RoleRepository struct{}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (rr *RoleRepository) Create(input *domain.Role) (*domain.Role, error) {
	db, _ := database.GetDatabaseConnection()
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (rr *RoleRepository) GetById(id int) (*domain.Role, error) {
	var role domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&role)

	checkRoleExist := db.Debug().Where("ID = ?", id)

	tx := checkRoleExist.Find(&role)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (rr *RoleRepository) GetByName(name string) (*domain.Role, error) {
	role := new(domain.Role)
	db, _ := database.GetDatabaseConnection()
	tx := db.Where("name = ?", name).First(&role)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return role, nil
}

func (a *RoleRepository) Update(input *domain.Role) (*domain.Role, error) {
	db, _ := database.GetDatabaseConnection()
	_, err := a.GetById(int(input.ID))
	if err != nil {
		return nil, errors.New("the model doesnt exists")
	}
	tx := db.Where("id = ?", input.ID).Save(input)
	if tx.Error != nil {
		return input, tx.Error
	}
	tx.Commit()
	return input, nil
}

func (a *RoleRepository) Get(id int) (*domain.Role, error) {
	var role domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&role)

	checkRoleExist := db.Debug().Where("ID = ?", id)

	tx := checkRoleExist.First(&role)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (a *RoleRepository) GetAll() (*[]domain.Role, error) {
	var roles []domain.Role
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&roles)

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
	deleted := db.Debug().Delete(role)
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
