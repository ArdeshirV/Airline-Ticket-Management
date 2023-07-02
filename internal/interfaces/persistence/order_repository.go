package persistence

import (
	"errors"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/database"
)

type OrderRepository interface {
	Create(input *domain.Order) (*domain.Order, error)
	Update(input *domain.Order) (*domain.Order, error)
	Get(id int) (*domain.Order, error)
	GetAll() (*[]domain.Order, error)
	Delete(id int) error
	GetItems(orderId uint) ([]domain.OrderItem, error)
	GetByOrderNum(orderNum string) (domain.Order, error)
}
type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r OrderRepositoryImpl) Create(input *domain.Order) (*domain.Order, error) {
	db, _ := database.GetDatabaseConnection() // todo: ignoring error, bad practice
	if input.ID > 0 {
		return nil, errors.New("can not create existing model")
	}
	db.Create(input)

	return input, nil
}

func (r OrderRepositoryImpl) Update(input *domain.Order) (*domain.Order, error) {
	db, _ := database.GetDatabaseConnection()
	_, err := r.Get(int(input.ID))
	println("status:", input.Status)
	if err != nil {
		return nil, errors.New("the model doesnt exists")
	}
	tx := db.Where("id = ?", input.ID).Save(input)
	if tx.Error != nil {
		return input, tx.Error
	}
	tx.Commit()
	return input, nil
}

func (r OrderRepositoryImpl) Get(id int) (*domain.Order, error) {
	var order domain.Order
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&order)

	tx := db.First(&order, id)
	if tx.Error != nil {
		return &order, tx.Error
	}
	return &order, nil
}

func (r OrderRepositoryImpl) GetAll() (*[]domain.Order, error) {
	var orders []domain.Order
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&orders)

	db.Find(&orders)

	return &orders, nil
}

func (r OrderRepositoryImpl) Delete(id int) error {
	order, err := r.Get(id)
	if err != nil {
		return err
	}
	db, _ := database.GetDatabaseConnection()
	db = db.Model(&order)
	tx := db.Delete(&order)
	if tx.RowsAffected > 0 {
		tx.Commit()
	}
	return nil
}

func (r OrderRepositoryImpl) GetItems(orderId uint) ([]domain.OrderItem, error) {
	var items []domain.OrderItem
	db, _ := database.GetDatabaseConnection()
	tx := db.Where("order_id = ?", orderId).Find(&items)
	if tx.Error != nil {
		return items, tx.Error
	}
	return items, nil
}
func (r OrderRepositoryImpl) GetByOrderNum(orderNum string) (domain.Order, error) {
	var order domain.Order
	db, _ := database.GetDatabaseConnection()
	tx := db.Where("order_num = ?", orderNum).First(&order)
	if tx.Error != nil {
		return order, tx.Error
	}
	return order, nil
}
