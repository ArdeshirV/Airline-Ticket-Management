package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
)

type LogoutHandler struct {
	usecase *usecase.UserUsecase
}

func (uh *UserHandler) Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get(config.Config.Auth.RequestLogoutHeader)
	// Check if the token is empty
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result: nil})
	}
	tokenString := strings.TrimPrefix(authHeader, config.Config.Auth.TokenPrefix+" ")
	JwtTokenSecretConfig := fmt.Sprintf("%v", config.Config.JwtToken.SecretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the key for verifying the token signature
		return []byte(JwtTokenSecretConfig), nil
	})

	if err != nil {
		println(err.Error())
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result: nil})
	}

	// check if token is valid
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result: nil})
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result: nil})
	}

	// Convert userId claim
	userId, _ := claims["userId"].(float64) // Todo: add to env ( it can be in env)
	intUserId := uint(userId)
	// select user from db
	user, err := uh.userUsecase.GetUserById(intUserId)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result: nil})
	}

	// update IsLoginRequired field in user
	user.IsLoginRequired = true
	uh.userUsecase.UpdateById(intUserId, user)

	return c.JSON(http.StatusOK, Response{Message: "you logged out successfully", Result: nil})
}
