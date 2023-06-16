package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/usecase"
)

type SignupRequest struct {
	Username string
	Email    string
	Password string
	Phone    string
}

type MassageResponse struct {
	Message string `json:"message"`
}

type SignupHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *SignupHandler {
	return &SignupHandler{
		usecase: usecase,
	}
}

func (sh *SignupHandler) Signup(c echo.Context) error {
	var request SignupRequest

	// Check the body data
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, MassageResponse{Message: "Invalid body request"})

	}
	if request.Username == "" || request.Email == "" || request.Password == "" || request.Phone == "" {
		return c.JSON(http.StatusBadRequest, MassageResponse{Message: "Missing required fields"})
	}

	// Check for dupplication
	_, err = sh.usecase.GetUserByEmail(request.Email)
	if err == nil {
		return c.JSON(http.StatusConflict, MassageResponse{Message: "User already exists with the given email or username"})
	}
	_, err = sh.usecase.GetUserByUsername(request.Username)
	if err == nil {
		return c.JSON(http.StatusConflict, MassageResponse{Message: "User already exists with the given email or username"})
	}

	// Create the user
	user := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		Phone:    request.Phone,
		RoleID: 1,
	}
	_, err = sh.usecase.CreateUser(&user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return c.JSON(http.StatusInternalServerError, MassageResponse{Message: "Cant create user"})
	}

	return c.JSON(http.StatusOK, MassageResponse{Message: "Created"})
}
