package repositories

import (
	"labireen/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entities.Order) error
	GetByID(param string, id uuid.UUID) (*entities.Order, error)
	GetWhere(param string, args string) (*entities.Order, error)
	Update(order *entities.Order) error
	Delete(order *entities.Order) error
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
	if err := r.db.Where(param+" = ?", id.String()).Preload("OrderStatuses.OrderItems").First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepositoryImpl) GetWhere(param string, args string) (*entities.Order, error) {
	var order entities.Order
	if err := r.db.Where(param+" = ?", args).Preload("OrderStatuses.OrderItems").First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepositoryImpl) Update(order *entities.Order) error {
	return r.db.Save(&order).Error
}

func (r *orderRepositoryImpl) Delete(order *entities.Order) error {
	return r.db.Delete(&order).Error
}
