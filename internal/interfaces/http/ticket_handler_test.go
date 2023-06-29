package http

import (
	"bytes"
	"encoding/json"
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
	ticketID := "123"
	req := httptest.NewRequest(http.MethodGet, "/ticket?ticketid="+ticketID, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	err := ticketHandler(ctx)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/pdf", rec.Header().Get("Content-Type"))
	assert.NotEmpty(t, rec.Body.Bytes())
	responseJSON := struct {
		Message string `json:"message"`
	}{}
	err = json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&responseJSON)
	assert.Nil(t, err)
	assert.Equal(t, "", responseJSON.Message)
}

func TestTicketHandler_InvalidTicketID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/ticket?ticketid=notanumber", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	err := ticketHandler(ctx)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get("Content-Type"))
	responseJSON := struct {
		message string `json:"error"`
	}{}
	err = json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&responseJSON)
	assert.Nil(t, err)
	assert.Equal(t, "failed to convert ticketid='notanumber' to integer", responseJSON.message)
}

func TestTicketHandler_MissingTicketID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/ticket", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	err := ticketHandler(ctx)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get("Content-Type"))
	responseJSON := struct {
		message string `json:"error"`
	}{}
	err = json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&responseJSON)
	assert.Nil(t, err)
	assert.Equal(t, "the 'ticketid' parameter is required", responseJSON.message)
}
