package http

import (
	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/pkg/config"
)

type APIResponse struct {
	Message string `json: "message"`
}

func (resp APIResponse) GetMessage() string {
	return resp.Message
}

func echoErrorAsJSON(ctx echo.Context, status int, err error) error {
	return echoJSON(ctx, status, APIResponse{Message: err.Error()})
}

func echoStringAsJSON(ctx echo.Context, status int, msg string) error {
	return echoJSON(ctx, status, APIResponse{Message: msg})
}

func echoJSON(ctx echo.Context, status int, data interface{}) error {
	if config.IsDebugMode() {
		return ctx.JSONPretty(status, data, "    ")
	} else {
		return ctx.JSON(status, data)
	}
}
