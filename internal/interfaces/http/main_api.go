package http

import (
	_ "fmt"
	_ "strconv"
	"net/http"
	"github.com/labstack/echo/v4"
	_ "github.com/the-go-dragons/final-project/internal/domain"
)

func MainRoutes(e *echo.Echo) {
	e.GET("/", listMainHandler)
	//e.POST("/", createMainHandler)  // Method samples
	//e.GET("/:id", findMainHandler)
	//e.DELETE("/:id", deleteMainHandler)
	//e.PUT("/", updateMainHandler)
}

func listMainHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "main_api.listMainHandler: TODO:data-goes-here")
}
