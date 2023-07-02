package http

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project/internal/usecase"
)

type PayError struct {
	Message string
}

type PayResult struct {
	Result   string
	OrederID int
}

type PaymentHandler struct {
	Payment *usecase.PaymentService
}

func NewPaymentHandler(payment *usecase.PaymentService) PaymentHandler {
	return PaymentHandler{Payment: payment}
}
func (p *PaymentHandler) Pay(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("orderId"))
	bank := c.QueryParam("bank")
	if bank == "" {
		return c.JSON(http.StatusBadRequest, PayError{Message: "Query parameter 'bank' is required"})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, PayError{Message: "OrderId should be integer"})
	}
	page, err := p.Payment.GetPaymentPage(orderId, bank)
	if err != nil {
		switch err.(type) {
		case usecase.OrderNotFound:
			return c.JSON(http.StatusNotFound, PayResult{Result: err.Error()})
		case usecase.InvalidBankName:
			return c.JSON(http.StatusBadRequest, PayResult{Result: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}
	return c.JSON(http.StatusOK, page)
}

func (p *PaymentHandler) Callback(c echo.Context) error {
	form, _ := c.FormParams()
	bank := c.QueryParam("bank")
	if bank == "" {
		return c.JSON(http.StatusBadRequest, PayError{Message: "Query parameter 'bank' is required"})
	}
	orderId, err := p.Payment.Callback(form, bank)
	if err != nil {
		switch err.(type) {
		case usecase.InvalidBankName:
			return c.JSON(http.StatusBadRequest, PayResult{Result: err.Error()})
		case usecase.VerifyingPaymentFailed:
			return c.JSON(http.StatusBadRequest, PayResult{Result: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}

	return c.JSON(http.StatusOK, PayResult{Result: "Successful", OrederID: orderId})

}
