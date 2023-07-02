package http

import (
	_ "bytes"
	_ "encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "github.com/the-go-dragons/final-project/internal/usecase"
	"github.com/the-go-dragons/final-project/pkg/config"
	_ "github.com/the-go-dragons/final-project/pkg/config"
	"github.com/the-go-dragons/final-project/pkg/database"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")
	config.Path = dir
}

func Setup() {
	config.Load()
	database.Load()
	database.CreateDBConnection()
}

func TestTicketHandler(t *testing.T) {
	Setup()
	e := echo.New()
	ticketID := "1"
	req := httptest.NewRequest(http.MethodGet, "/print_ticket?ticketid="+ticketID, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	err := ticketHandler(ctx)
	assert.Error(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
