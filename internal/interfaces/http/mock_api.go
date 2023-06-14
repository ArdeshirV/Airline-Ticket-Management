package http

import (
	_ "fmt"
	_ "strconv"
	"net/http"
	"github.com/labstack/echo/v4"
	_ "github.com/the-go-dragons/final-project/internal/domain"
)

func MockRoutes(e *echo.Echo) {
	e.GET("/mock", listMockHandler)
	//e.POST("/mock", createMockHandler)  // Method samples
	//e.GET("/mock/:id", findMockHandler)
	//e.DELETE("/mock/:id", deleteMockHandler)
	//e.PUT("/mock", updateMockHandler)
}

func listMockHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "TODO:data-goes-here")
}
