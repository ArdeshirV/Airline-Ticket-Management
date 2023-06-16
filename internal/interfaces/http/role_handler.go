package http

import "github.com/the-go-dragons/final-project/internal/usecase"

type RoleHandler struct {
	roleUsecase *usecase.RoleUsecase
}

func NewRoleHandler(usecase *usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		roleUsecase: usecase,
	}
}
