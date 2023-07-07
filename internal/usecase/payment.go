package usecase

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type PaymentService struct {
	paymentRepo persistence.PaymentRepository
	orderRepo   persistence.OrderRepository
}

func (PaymentService) GetGateway(bank Bank) Gateway {
	return BANKS[bank]()
}
func NewPayment(
	paymentRepo *persistence.PaymentRepository,
	orderRepo persistence.OrderRepository,
) *PaymentService {
	return &PaymentService{paymentRepo: *paymentRepo, orderRepo: orderRepo}
}

func (p PaymentService) GetPaymentPage(orderID int, bankName string) (PaymentPage, error) {
	bank, err := validateAndGetBank(bankName)
	if err != nil {
		return PaymentPage{}, InvalidBankName{bankName}
	}
	order, err := p.orderRepo.Get(orderID)
	if err != nil {
		return PaymentPage{}, OrderNotFound{orderID}
	}
	token := p.GetGateway(bank).GetToken(*order)
	url := p.GetGateway(bank).GetPaymentPage(token)
	return url, nil
}

func (p PaymentService) Callback(data map[string][]string, bankName string) (int, error) {
	bank, err := validateAndGetBank(bankName)
	if err != nil {
		return 0, InvalidBankName{bankName}
	}
	payment, err := p.GetGateway(bank).VerifyPayment(data)
	if err != nil {
		return 0, VerifyingPaymentFailed{int(payment.OrderID)}
	}

	// order, err := p.orderRepo.GetByOrderNum(orderNum)

	order, err := p.orderRepo.Get(int(payment.OrderID))
	if err != nil {
		return 0, err
	}

	if payment.PayAmount != order.Amount {
		return 0, errors.New("wrong payment amount for order")
	}
	p.paymentRepo.Create(&payment)
	order.Status = domain.PAID
	p.orderRepo.Update(order)
	return int(order.ID), nil
}
