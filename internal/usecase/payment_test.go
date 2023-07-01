package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/the-go-dragons/final-project/internal/domain"
)

type GatewayMock struct {
	mock.Mock
}

func (s GatewayMock) GetName() string {
	return ""
}

func (s GatewayMock) GetPaymentPage(token string) PaymentPage {
	args := s.Called(token)
	return args.Get(0).(PaymentPage)
}

func (s GatewayMock) VerifyPayment(data map[string][]string) (domain.Payment, error) {
	args := s.Called(data)
	return args.Get(0).(domain.Payment), args.Error(1)
}

func (s GatewayMock) GetToken(order domain.Order) string {
	args := s.Called(order)
	return args.String(0)
}

func TestGetPaymentPage(t *testing.T) {
	repo := new(OrderRepoMock)
	gateway := new(GatewayMock)
	orderId := 1
	bankName := "saderat"
	token := "0123"
	page := PaymentPage{URL: "http://test.com"}
	order := domain.Order{OrderNum: "Test"}
	order.ID = uint(orderId)
	repo.On("Get", orderId).Return(&order, nil)
	gateway.On("GetToken", order).Return(token)
	gateway.On("GetPaymentPage", token).Return(page)
	BANKS[Saderat] = func() Gateway { return gateway }

	payment := PaymentService{orderRepo: repo}
	result, err := payment.GetPaymentPage(orderId, bankName)

	assert.Equal(t, page.URL, result.URL)
	assert.Nil(t, err)
}

func TestGetPaymentPageInvalidBankName(t *testing.T) {
	repo := new(OrderRepoMock)
	gateway := new(GatewayMock)
	orderId := 1
	bankName := "sderat"
	token := "0123"
	page := PaymentPage{URL: "http://test.com"}
	order := domain.Order{OrderNum: "Test"}
	order.ID = uint(orderId)
	repo.On("Get", orderId).Return(&order, nil)
	gateway.On("GetToken", order).Return(token)
	gateway.On("GetPaymentPage", token).Return(page)
	BANKS[Saderat] = func() Gateway { return gateway }

	payment := PaymentService{orderRepo: repo}
	_, err := payment.GetPaymentPage(orderId, bankName)
	assert.EqualError(t, err, InvalidBankName{bankName}.Error())
}

func TestGetPaymentPageOrderNotFound(t *testing.T) {
	repo := new(OrderRepoMock)
	orderId := 1
	bankName := "saderat"
	order := domain.Order{OrderNum: "Test"}
	errs := errors.New("error")
	repo.On("Get", orderId).Return(&order, errs)

	payment := PaymentService{orderRepo: repo}
	_, err := payment.GetPaymentPage(orderId, bankName)

	assert.EqualError(t, err, OrderNotFound{orderId}.Error())
}

func TestCallback(t *testing.T) {
	repo := new(OrderRepoMock)
	payRepo := new(PaymentRepoMock)
	gateway := new(GatewayMock)
	data := map[string][]string{}
	orderId := 1
	bankName := "saderat"
	payment := domain.Payment{
		OrderID:   uint(orderId),
		PayAmount: 5000,
	}
	order := domain.Order{
		OrderNum: "Test",
		Amount:   5000,
	}
	order.ID = uint(orderId)
	gateway.On("VerifyPayment", data).Return(payment, nil)
	repo.On("Get", orderId).Return(&order, nil)
	BANKS[Saderat] = func() Gateway { return gateway }

	paymentService := PaymentService{
		orderRepo:   repo,
		paymentRepo: payRepo,
	}
	result, err := paymentService.Callback(data, bankName)

	assert.Equal(t, orderId, result)
	assert.Nil(t, err)
}

func TestCallbackInvalidBankName(t *testing.T) {
	repo := new(OrderRepoMock)
	payRepo := new(PaymentRepoMock)
	gateway := new(GatewayMock)
	data := map[string][]string{}
	orderId := 1
	bankName := "sderat"
	payment := domain.Payment{
		OrderID:   uint(orderId),
		PayAmount: 5000,
	}
	order := domain.Order{
		OrderNum: "Test",
		Amount:   5000,
	}
	order.ID = uint(orderId)
	gateway.On("VerifyPayment", data).Return(payment, nil)
	errs := errors.New("error")
	repo.On("Get", orderId).Return(&order, errs)
	BANKS[Saderat] = func() Gateway { return gateway }

	paymentService := PaymentService{
		orderRepo:   repo,
		paymentRepo: payRepo,
	}
	_, err := paymentService.Callback(data, bankName)

	assert.EqualError(t, err, InvalidBankName{bankName}.Error())
}

func TestCallbackVerifyingPaymentFailed(t *testing.T) {
	repo := new(OrderRepoMock)
	payRepo := new(PaymentRepoMock)
	gateway := new(GatewayMock)
	data := map[string][]string{}
	orderId := 1
	bankName := "saderat"
	payment := domain.Payment{
		OrderID:   uint(orderId),
		PayAmount: 5000,
	}
	order := domain.Order{
		OrderNum: "Test",
		Amount:   5000,
	}
	order.ID = uint(10)
	errs := errors.New("error")
	gateway.On("VerifyPayment", data).Return(payment, errs)
	repo.On("Get", orderId).Return(&order, nil)
	BANKS[Saderat] = func() Gateway { return gateway }

	paymentService := PaymentService{
		orderRepo:   repo,
		paymentRepo: payRepo,
	}
	_, err := paymentService.Callback(data, bankName)

	assert.EqualError(t, err, VerifyingPaymentFailed{orderId}.Error())
}

func TestCallbackWrongAmount(t *testing.T) {
	repo := new(OrderRepoMock)
	payRepo := new(PaymentRepoMock)
	gateway := new(GatewayMock)
	data := map[string][]string{}
	orderId := 1
	bankName := "saderat"
	payment := domain.Payment{
		OrderID:   uint(orderId),
		PayAmount: 1000,
	}
	order := domain.Order{
		OrderNum: "Test",
		Amount:   5000,
	}
	order.ID = uint(1)
	gateway.On("VerifyPayment", data).Return(payment, nil)
	repo.On("Get", orderId).Return(&order, nil)
	BANKS[Saderat] = func() Gateway { return gateway }

	paymentService := PaymentService{
		orderRepo:   repo,
		paymentRepo: payRepo,
	}
	_, err := paymentService.Callback(data, bankName)

	assert.EqualError(t, err, "wrong payment amount for order")
}
