package repositories

import (
	"labireen/customer_microservices/account_service/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateCustomer(customer *entities.Customer) error
	GetCustomerById(id uuid.UUID) (*entities.Customer, error)
	GetCustomerByEmail(email string) (*entities.Customer, error)
	GetCustomerByCustom(conds string, args string) (*entities.Customer, error)
	UpdateCustomer(customer *entities.Customer) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) CreateCustomer(customer *entities.Customer) error {
	return r.db.Create(&customer).Error
}

func (r *authRepositoryImpl) GetCustomerById(id uuid.UUID) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.First(customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *authRepositoryImpl) GetCustomerByEmail(email string) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.Where("email = ?", email).Take(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *authRepositoryImpl) GetCustomerByCustom(conds string, args string) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.Where(conds+" = ?", args).Take(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *authRepositoryImpl) UpdateCustomer(customer *entities.Customer) error {
	return r.db.Save(customer).Error
}
