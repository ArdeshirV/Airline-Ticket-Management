package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	_ "github.com/the-go-dragons/final-project/internal/domain"
)

func TestAddPassenger(t *testing.T) {
	e := echo.New()

	requestBody := map[string]interface{}{
		"firstName":    "Akbar",
		"lastName":     "Akbari",
		"nationalCode": "1234567890",
		"email":        "akbar.akabari@example.com",
		"gender":       "male",
		"birthDate":    "1990-01-01",
		"phone":        "1234567890",
		"address":      "123 Example Street",
	}
	requestJSON, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/passengers", bytes.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := AddPassenger(c)
	assert.NoError(t, err)
}
