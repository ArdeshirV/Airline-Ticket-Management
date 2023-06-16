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

func (rr *RoleRepository) Create(role *domain.Role) (*domain.Role, error) {
	
	db, _ := database.GetDatabaseConnection()
	result := db.Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}

func (rr *RoleRepository) GetById(id uint) (*domain.Role, error) {
	role := new(domain.Role)
	db, _ := database.GetDatabaseConnection()
	db.Where("id = ?", id).First(&role)
	if role.ID == 0 {
		return nil, errors.New("Role not found")
	}
	return role, nil
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