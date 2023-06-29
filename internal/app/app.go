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
	secret = config.Config.JwtToken.ExpireHours
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

	flightRepo := persistence.NewFlightRepository()
	passengerRepo := persistence.NewPassengerRepository()
	orderRepo := persistence.NewOrderRepository()
	paymentRepo := persistence.NewPaymentRepository()

	payment := usecase.NewPayment(paymentRepo, orderRepo)
	PaymentHandler := handlers.NewPaymentHandler(payment)

	ticketRepo := persistence.NewTicketRepository()
	ticketUsecase := usecase.NewTicketUsecase(ticketRepo)

	booking := usecase.NewBooking(flightRepo, passengerRepo, orderRepo, ticketRepo)
	BookingHandler := handlers.NewBookingHandler(booking)
	flightUseCase := usecase.NewFlightUseCase(flightRepo)
	TicketHandler := handlers.NewTicketHandler(ticketUsecase, flightUseCase, booking)

	// UserHandler := handlers.NewUserHandler(roleUsecase)
	_ = RoleHandler

	// public routing
	handlers.RootRoute(e)
	handlers.FlightsRoute(e)
	handlers.PassengerRoute(e)
	handlers.PrintTicketRoute(e)

	e.POST("/signup", UserHandler.Signup)
	e.POST("/login", UserHandler.Login)
	e.GET("/logout", UserHandler.Logout)
	e.GET("/payment/pay/:orderId", PaymentHandler.Pay)
	e.POST("/payment/callback", PaymentHandler.Callback)
	e.POST("/booking/book", BookingHandler.Book)
	e.POST("/booking/finalize", BookingHandler.Finalize)

	e.GET("/reserved", TicketHandler.GetReservedUsers)
	e.POST("/cancel", TicketHandler.Cancel)
}

func initializeSessionStore() {
	store = sessions.NewCookieStore([]byte(fmt.Sprintf("%v", secret)))

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
