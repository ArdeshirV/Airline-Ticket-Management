package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	html = `<strong style="color: green; background-color:black; text-align:center; ">Final Project of The Go Dragons Team</strong>
<br/>
<a style="text-alighn:center; " target="_blank" rel="noopener noreferrer" href="https://github.com/the-go-dragons/final-project">Visit Project on Github</a>`
)

func RootRoute(e *echo.Echo) {
	e.GET("/", root)
}

func root(c echo.Context) error {
	return c.HTML(http.StatusOK, html)
}
