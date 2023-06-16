package usecase

import "github.com/the-go-dragons/final-project/internal/interfaces/persistence"

type RoleUsecase struct {
	repository *persistence.RoleRepository
}

func NewRoleUsecase(repository *persistence.RoleRepository) *RoleUsecase {
	return &RoleUsecase{
		repository: repository,
	}
}