package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	handlers "github.com/the-go-dragons/final-project/internal/interfaces/http"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"github.com/the-go-dragons/final-project/internal/usecase"
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
	userRepo := persistence.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	UserHandler := handlers.NewUserHandler(userUsecase)

	handlers.MockRoutes(e)
	handlers.MainRoutes(e)

	// public routing
	e.POST("/signup", UserHandler.Signup)
	e.POST("/login", UserHandler.Login)
	// e.POST("/token", UserController.GetToken)

	// protected routing
	// e.GET("/now", UserController.GetTime, middlewares.IsLoggedIn, middlewares.IsAdmin)
}
