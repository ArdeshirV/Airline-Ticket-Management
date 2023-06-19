package app

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlers "github.com/the-go-dragons/final-project/internal/interfaces/http"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
)

var (
	store  = sessions.NewCookieStore()
	secret = config.GetEnv("JWT_TOKEN_EXPIRE_HOURS", "mySecretKey")
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

	initializeSessionStore()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(SessionMiddleware())

	userRepo := persistence.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)

	roleRepo := persistence.NewRoleRepository()
	roleUsecase := usecase.NewRoleUsecase(roleRepo)

	UserHandler := handlers.NewUserHandler(userUsecase, roleUsecase)
	RoleHandler := handlers.NewRoleHandler(roleUsecase)

	// UserHandler := handlers.NewUserHandler(roleUsecase)

	handlers.MockRoutes(e)
	handlers.MainRoutes(e)

	_ = RoleHandler

	// public routing
	e.POST("/signup", UserHandler.Signup)
	e.POST("/login", UserHandler.Login)
	e.GET("/logout", UserHandler.Logout)

	//e.GET("/protected", SomeProtectedRouteHandler, UserHandler.Authorize)

	// protected routing
	// e.GET("/now", UserController.GetTime, middlewares.IsLoggedIn, middlewares.IsAdmin)
}

func initializeSessionStore() {
	store = sessions.NewCookieStore([]byte(secret))

	// Set session options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // Session expiration time (in seconds)
		HttpOnly: true,
	}
}

func SessionMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			session, _ := store.Get(c.Request(), "go-dragon-session")
			c.Set("session", session)

			return next(c)
		}
	}
}

/*func SomeProtectedRouteHandler(c echo.Context) error {
	println("salam")

	return nil
}*/
