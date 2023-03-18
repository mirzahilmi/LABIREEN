package repositories

import (
	"labireen/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entities.Order) error
	GetByID(param string, id uuid.UUID) (*entities.Order, error)
	GetAllByCustomer(id uuid.UUID) (*[]entities.Order, error)
	GetAllByMerchant(id uuid.UUID) (*[]entities.Order, error)
	Update(order *entities.Order) error
	Delete(id uuid.UUID) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (r *orderRepositoryImpl) Create(order *entities.Order) error {
	return r.db.Create(&order).Error
}

func (r *orderRepositoryImpl) GetByID(param string, id uuid.UUID) (*entities.Order, error) {
	var order entities.Order
	if err := r.db.Where(param+" = ?", id).Preload("OrderStatuses.OrderItems").First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepositoryImpl) GetAllByCustomer(id uuid.UUID) (*[]entities.Order, error) {
	var orders []entities.Order
	if err := r.db.Where("customer_id = ?", id).Preload("OrderStatuses.OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}

	return &orders, nil
}
func (r *orderRepositoryImpl) GetAllByMerchant(id uuid.UUID) (*[]entities.Order, error) {
	var orders []entities.Order
	if err := r.db.Where("merchant_id = ?", id).Preload("OrderStatuses.OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r *orderRepositoryImpl) Update(order *entities.Order) error {
	return r.db.Save(&order).Error
}

func (r *orderRepositoryImpl) Delete(id uuid.UUID) error {
	var orders []entities.Order
	if err := r.db.Where("merchant_id = ?", id).Preload("OrderStatuses.OrderItems").Find(&orders).Error; err != nil {
		return err
	}

	return nil
}
