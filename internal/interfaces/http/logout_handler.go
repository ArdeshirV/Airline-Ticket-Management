package http

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
)

type LogoutHandler struct {
	usecase *usecase.UserUsecase
}

func (uh *UserHandler) Logout(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization") // TODO: set in env
	// Check if the token is empty
    if tokenString == "" {
        // Return a 401 Unauthorized response if no token was provided
        return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result:nil})

	}
	
	JwtTokenSecretConfig := config.GetEnv("JWT_TOKEN_EXPIRE_HOURS", "mySecretKey")

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Return the key for verifying the token signature
        return []byte(JwtTokenSecretConfig), nil
    })

	// Check if the token is valid
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result:nil})
	} 

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result:nil})
	}

	userId, _ := claims["userId"].(float64)
	fmt.Printf("userId: %v\n", userId)
	intUserId := uint(userId)
	fmt.Printf("intUserId: %v\n", intUserId)

	user, error := uh.userUsecase.GetUserById(intUserId)
	
	if error != nil {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Authoization header is not valid", Result:nil})
	}

	user.IsLoginRequired = true
	uh.userUsecase.UpdateById(intUserId, user)

	return c.JSON(http.StatusOK, Response{Message: "you logged out successfully", Result:nil })

	return nil
}