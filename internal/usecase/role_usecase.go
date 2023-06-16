package usecase

import (
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type RoleUsecase struct {
	roleRepository *persistence.RoleRepository
}

func NewRoleUsecase(repository *persistence.RoleRepository) *RoleUsecase {
	return &RoleUsecase{
		roleRepository: repository,
	}
}

func (ru *RoleUsecase) CreateRole(role *domain.Role) (*domain.Role, error) {
	return ru.roleRepository.Create(role)
}

func (ru *RoleUsecase) GetById(id uint) (*domain.Role, error) {
	return ru.roleRepository.GetById(id)
}

func (ru *RoleUsecase) GetByName(name string) (*domain.Role, error) {
	return ru.roleRepository.GetByName(name)
}