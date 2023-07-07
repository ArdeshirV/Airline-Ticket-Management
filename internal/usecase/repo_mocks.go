package usecase

import (
	"github.com/stretchr/testify/mock"
	"github.com/the-go-dragons/final-project/internal/domain"
)

type OrderRepoMock struct {
	mock.Mock
}

func (r OrderRepoMock) Create(input *domain.Order) (*domain.Order, error) {
	return nil, nil
}

func (r OrderRepoMock) Update(input *domain.Order) (*domain.Order, error) {
	return nil, nil
}

func (r OrderRepoMock) Get(id int) (*domain.Order, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (r OrderRepoMock) GetAll() (*[]domain.Order, error) {
	return nil, nil
}

func (r OrderRepoMock) Delete(id int) error {
	return nil
}

func (r OrderRepoMock) GetItems(orderId uint) ([]domain.OrderItem, error) {
	return make([]domain.OrderItem, 0), nil
}
func (r OrderRepoMock) GetByOrderNum(orderNum string) (domain.Order, error) {
	return domain.Order{}, nil
}

type PaymentRepoMock struct {
	mock.Mock
}

func (p PaymentRepoMock) Create(input *domain.Payment) (*domain.Payment, error) {
	return nil, nil
}
func (p PaymentRepoMock) Update(input *domain.Payment) (*domain.Payment, error) {
	return nil, nil
}
func (p PaymentRepoMock) Get(id int) (*domain.Payment, error) {
	return nil, nil
}
func (p PaymentRepoMock) GetByOrderId(orderID int) (*domain.Payment, error) {
	return nil, nil
}

type FlightRepoMock struct {
	mock.Mock
}

func (f FlightRepoMock) Create(input *domain.Flight) (*domain.Flight, error) {
	return nil, nil
}
func (f FlightRepoMock) Update(input *domain.Flight) (*domain.Flight, error) {
	return nil, nil
}
func (f FlightRepoMock) Get(id int) (*domain.Flight, error) {
	args := f.Called(id)
	return args.Get(0).(*domain.Flight), args.Error(1)
}
func (f FlightRepoMock) GetAll() (*[]domain.Flight, error) {
	return nil, nil
}
func (f FlightRepoMock) Delete(id int) error {
	return nil
}

type PassengerRepoMock struct {
	mock.Mock
}

func (p PassengerRepoMock) Create(input *domain.Passenger) (*domain.Passenger, error) {
	return nil, nil
}
func (p PassengerRepoMock) Update(input *domain.Passenger) (*domain.Passenger, error) {
	return nil, nil
}
func (p PassengerRepoMock) Get(id int) (*domain.Passenger, error) {
	return nil, nil
}
func (p PassengerRepoMock) GetAll() (*[]domain.Passenger, error) {
	return nil, nil
}
func (p PassengerRepoMock) GetList(IDs []int) ([]domain.Passenger, error) {
	return make([]domain.Passenger, 0), nil
}
func (p PassengerRepoMock) Delete(id int) error {
	return nil
}
