package usecase

import (
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/internal/interfaces/persistence"
)

type Payment struct {
	paymentRepo persistence.PaymentRepository
	orderRepo   persistence.OrderRepository
}

func (Payment) GetGateway(bank Bank) Gateway {
	return BANKS[bank]
}
func NewPayment(
	paymentRepo *persistence.PaymentRepository,
	orderRepo *persistence.OrderRepository,
) *Payment {
	return &Payment{paymentRepo: *paymentRepo, orderRepo: *orderRepo}
}

func (p Payment) GetPaymentPage(orderID int, bankName string) (PaymentPage, error) {
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

func (p Payment) Callback(data map[string][]string, bankName string) (bool, error) {
	bank, err := validateAndGetBank(bankName)
	if err != nil {
		return false, InvalidBankName{bankName}
	}
	payment, err := p.GetGateway(bank).VerifyPayment(data)
	if err != nil {
		return false, nil
	}
	orderNum := data["invoiceid"][0]
	// order, err := p.orderRepo.GetByOrderNum(orderNum)
	id, _ := strconv.Atoi(orderNum)
	order, err := p.orderRepo.Get(id)
	if err != nil {
		return false, nil
	}
	payment.OrderID = order.ID
	if payment.PayAmount != order.Amount {
		return false, nil
	}
	p.paymentRepo.Create(&payment)
	order.Status = domain.PAID
	p.orderRepo.Update(order)
	return true, nil
}
