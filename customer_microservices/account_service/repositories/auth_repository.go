package repositories

import (
	"labireen/customer_microservices/account_service/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(customer *entities.Customer) error
	GetById(id uuid.UUID) (*entities.Customer, error)
	GetWhere(param string, email string) (*entities.Customer, error)
	Update(customer *entities.Customer) error
	Delete(customer *entities.Customer) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) Create(customer *entities.Customer) error {
	return r.db.Create(&customer).Error
}

func (r *authRepositoryImpl) GetById(id uuid.UUID) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *authRepositoryImpl) GetWhere(param string, email string) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.Where(param+" = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *authRepositoryImpl) Update(customer *entities.Customer) error {
	return r.db.Save(customer).Error
}

func (r *authRepositoryImpl) Delete(customer *entities.Customer) error {
	return r.db.Delete(customer).Error
}
