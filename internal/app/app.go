package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	handlers "github.com/the-go-dragons/final-project/internal/interfaces/http"
)

type App struct {
	E *echo.Echo
}

func NewApp() *App {
	e := echo.New()
	routing(e)

	return &App{
		E: e,
	}
}

func (application *App) Start(portAddress string) error {
	fmt.Println("portAddress =", portAddress)
	err := application.E.Start(fmt.Sprintf(":%s", portAddress))
	application.E.Logger.Fatal(err)
	return err
}

func routing(e *echo.Echo) {
	// userRepo := userRepository.NewGormUserRepository()
	// UserService := userService.NewUserService(userRepo)
	// UserController := usercontroller.UserController{UserService: UserService}

	handlers.MockRoutes(e)
	handlers.MainRoutes(e)

	// public routing
	// e.POST("/signup", UserController.Signup)
	// e.POST("/login", UserController.Login)
	// e.POST("/token", UserController.GetToken)

	// protected routing
	// e.GET("/now", UserController.GetTime, middlewares.IsLoggedIn, middlewares.IsAdmin)
}
