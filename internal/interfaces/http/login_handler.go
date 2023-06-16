package http

import (
	"fmt"
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

type LoginMassageResponse struct {
	Message string `json:"message"`
	Result *LoginResult `json:"result"`
}

type LoginHandler struct {
	usecase *usecase.UserUsecase
}

func GenerateToken (user *domain.User) (string, error) {
	expirationHoursCofig := config.GetEnv("JWT_TOKEN_EXPIRE_HOURS", "1")
	JwtTokenSecretConfig := config.GetEnv("JWT_TOKEN_EXPIRE_HOURS", "mySecretKey")

	expirationCofigHoursValue, _ := strconv.ParseUint(expirationHoursCofig, 10, 64)
	uintExpirationCofigHoursValue := uint(expirationCofigHoursValue)

	duration := time.Duration(uintExpirationCofigHoursValue) * time.Hour
	expirationTime := time.Now().Add(duration) 

	idBytes := []byte(fmt.Sprintf("%d", user.ID))
	hashedUserId, err := bcrypt.GenerateFromPassword(
		idBytes,
		bcrypt.DefaultCost,
	)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": hashedUserId,
		"exp":    expirationTime.Unix(),
	})

	secretKey := []byte(JwtTokenSecretConfig)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (sh *SignupHandler) Login(c echo.Context) error {
	var request LoginRequest
	var user *domain.User

	// Check the body data
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginMassageResponse{Message: "Invalid body request", Result:nil})
	}

	if  request.Email == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, LoginMassageResponse{Message: "Missing required fields", Result:nil})
	}

	// Check for dupplication
	user, err = sh.usecase.GetUserByEmail(request.Email)
	if err != nil {
		return c.JSON(http.StatusConflict, LoginMassageResponse{Message: "No user found with this credentials", Result:nil})
	}

	equalErr := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password))
	
	if equalErr == nil {
		token, err := GenerateToken(user)
		fmt.Printf("\"here\": %v\n", "here")
		if err != nil {
			return c.JSON(http.StatusBadRequest, LoginMassageResponse{Message: "Server Error", Result:nil})
		}

		result := &LoginResult{
			Token: token,
		}

		return c.JSON(http.StatusOK, LoginMassageResponse{Message: "login", Result:result})
	}

	return c.JSON(http.StatusConflict, LoginMassageResponse{Message: "No user found with this credentials", Result:nil })
}