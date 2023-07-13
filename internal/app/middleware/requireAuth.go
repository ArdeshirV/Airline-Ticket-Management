package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type MassageResponse struct {
	Message string `json:"message"`
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Must authenticate"})
		}

		// Decode/validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(config.Config.JwtToken.SecretKey), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: fmt.Sprint("Invalid token: ", err.Error())})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// check the expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Token expired"})
			}

			// find the user
			var user domain.User
			db, _ := database.GetDatabaseConnection()
			db.First(&user, claims["userId"])

			if user.ID == 0 {
				return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "User not found"})
			}

			if user.IsLoginRequired {
				return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "User need to login"})
			}

			c.Set("user", user)
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Invalid token"})
	}

}
