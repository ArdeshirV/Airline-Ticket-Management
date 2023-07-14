package app

import (
	"fmt"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlers "github.com/the-go-dragons/final-project/internal/interfaces/http"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
	"github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
)

var (
	store     = sessions.NewCookieStore()
	getSecret = func() string { return config.Config.JwtToken.SecretKey }
)

type App struct {
	e *echo.Echo
}

func NewApp() *App {
	e := echo.New()
	routing(e)

	return &App{
		e: e,
	}
}

func (application *App) Start(portAddress int) error {
	application.e.Server.ReadTimeout = 5 * time.Second
	application.e.Server.WriteTimeout = 5 * time.Second
	err := application.e.Start(fmt.Sprintf(":%d", portAddress))
	application.e.Logger.Fatal(err)
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

	flightRepo := persistence.NewFlightRepository()
	passengerRepo := persistence.NewPassengerRepository()
	orderRepo := persistence.NewOrderRepository()
	paymentRepo := persistence.NewPaymentRepository()

	ticketRepo := persistence.NewTicketRepository()
	ticketUsecase := usecase.NewTicketUsecase(ticketRepo)

	payment := usecase.NewPayment(&paymentRepo, orderRepo)
	PaymentHandler := handlers.NewPaymentHandler(payment)

	booking := usecase.NewBooking(flightRepo, passengerRepo, orderRepo, paymentRepo)
	BookingHandler := handlers.NewBookingHandler(booking, *UserHandler)

	flightUseCase := usecase.NewFlightUseCase(flightRepo)

	TicketHandler := handlers.NewTicketHandler(ticketUsecase, flightUseCase, booking, *UserHandler)
	passenegerRepo := persistence.NewPassengerRepository()
	passengerHandler := handlers.NewPassegerHandler(passenegerRepo, *UserHandler)

	// UserHandler := handlers.NewUserHandler(roleUsecase)
	_ = RoleHandler

	// public routing
	handlers.RootRoute(e)
	handlers.DataRoute(e)
	handlers.FlightsRoute(e)
	handlers.PassengerRoute(passengerHandler, e)
	handlers.PrintTicketRoute(e)

	e.POST("/signup", UserHandler.Signup)
	e.POST("/login", UserHandler.Login)
	e.GET("/logout", UserHandler.Logout)
	e.GET("/payment/pay/:orderId", PaymentHandler.Pay)
	e.POST("/payment/callback", PaymentHandler.Callback)
	e.POST("/booking/book", BookingHandler.Book)
	e.POST("/booking/finalize", BookingHandler.Finalize)

	e.GET("/reserved", TicketHandler.GetReservedUsers /*, customMiddleware.RequireAuth*/)
	e.POST("/cancel", TicketHandler.Cancel)
}

func initializeSessionStore() {
	store = sessions.NewCookieStore([]byte(getSecret()))

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
