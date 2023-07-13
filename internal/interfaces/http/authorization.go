package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/domain"
)

func protectedEndpoint(c echo.Context) error {
	return c.String(http.StatusOK, "Protected endpoint")
}

// Handler function for the public endpoint
func publicEndpoint(c echo.Context) error {
	return c.String(http.StatusOK, "Public endpoint")
}

func NewAuthMiddleware(user *domain.User) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from the request header or query parameter
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				token = c.QueryParam("token")
			}

			if user.IsLoginRequired {
				return c.JSON(http.StatusNetworkAuthenticationRequired, map[string]string{"error": "Authentication Required!"})
			} else if token != "valid_token" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			SetUserToSession(c, user)
			// Call the next handler in the chain
			return next(c)
		}
	}
}

func (u *UserHandler) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := u.GetUserFromSession(c)
		if err != nil {
			return err
		}

		// Check if the user has the required role
		if !isAuthorized(user, "admin") {
			return echo.NewHTTPError(http.StatusForbidden, "Unauthorized")
		}

		// User is authorized
		return next(c)
	}
}

func (u *UserHandler) GetUserFromSession(c echo.Context) (*domain.User, error) {
	session := c.Get("session").(*sessions.Session)
	key := "userID"
	if userID, ok := session.Values[key]; ok {
		user, err := u.userUsecase.GetUserById(userID.(uint))
		if err != nil {
			return nil, err
		}
		println("inside 4")
		return user, nil
	}

	return nil, errors.New("No user found")
}

func isAuthorized(user *domain.User, role string) bool {
	return strings.EqualFold(strings.ToLower(user.Role.Name), strings.ToLower(role))
}
