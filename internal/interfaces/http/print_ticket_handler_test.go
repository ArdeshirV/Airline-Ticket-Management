package http

import (
	_ "bytes"
	_ "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "github.com/the-go-dragons/final-project/internal/usecase"
	_ "github.com/the-go-dragons/final-project/pkg/config"
)

func TestTicketHandler(t *testing.T) {
	e := echo.New()
	ticketID := "1"
	req := httptest.NewRequest(http.MethodGet, "/print_ticket?ticketid="+ticketID, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	err := ticketHandler(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
