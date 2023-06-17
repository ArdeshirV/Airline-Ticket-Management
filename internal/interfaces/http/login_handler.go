package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResult struct {
	Token string `json:"token"`
}

type Response struct {
	Message string       `json:"message"`
	Result  *LoginResult `json:"result"`
}

type LoginHandler struct {
	usecase *usecase.UserUsecase
}

func GenerateToken(user *domain.User) (string, error) {
	expirationHoursCofig := config.Get(config.JwtTokenExpireHours)
	JwtTokenSecretConfig := config.Get(config.JwtTokenSecretKey)

	expirationCofigHoursValue, _ := strconv.ParseUint(expirationHoursCofig, 10, 64)
	uintExpirationCofigHoursValue := uint(expirationCofigHoursValue)

	duration := time.Duration(uintExpirationCofigHoursValue) * time.Hour
	expirationTime := time.Now().Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    expirationTime.Unix(),
	})

	secretKey := []byte(JwtTokenSecretConfig)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uh *UserHandler) Login(c echo.Context) error {
	var request LoginRequest
	var user *domain.User

	// Check the body data
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: "Invalid body request", Result: nil})
		// TODO: all Responses should be in a standard

		// TODO: response messages should be mutli language
		// we can use i18 library
	}

	if request.Email == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, Response{Message: "Missing required fields", Result: nil})
	}

	// Check for dupplication
	user, err = uh.userUsecase.GetUserByEmail(request.Email)
	if err != nil {
		return c.JSON(http.StatusConflict, Response{Message: "No user found with this credentials", Result: nil})
	}

	// Check if password is correct
	equalErr := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password))

	if equalErr == nil {
		token, err := GenerateToken(user)

		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{Message: "Server Error", Result: nil})
		}

		result := &LoginResult{
			Token: token,
		}

		// update IsLoginRequired field
		user.IsLoginRequired = false
		uh.userUsecase.UpdateById(uint(user.ID), user)

		return c.JSON(http.StatusOK, Response{Message: "You logged in successfully", Result: result})
	}

	return c.JSON(http.StatusConflict, Response{Message: "No user found with this credentials", Result: nil})
}
