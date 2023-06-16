package http

import "github.com/the-go-dragons/final-project/internal/usecase"

func NewRoleHandler(usecase *usecase.UserUsecase) *SignupHandler {
	return &SignupHandler{
		usecase: usecase,
	}
}