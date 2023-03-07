package services

import (
	"labireen/customer_microservices/account_service/entities"
	"labireen/customer_microservices/account_service/repositories"
)

type CustomerService interface {
	UpdateCustomer(customer entities.CustomerRequest) error
	GetCustomer(email string) (entities.CustomerRequest, error)
}

type customerServiceImpl struct {
	repo repositories.AuthRepository
}

func NewCustomerService(repo repositories.AuthRepository) CustomerService {
	return &customerServiceImpl{repo}
}

func (csr *customerServiceImpl) UpdateCustomer(customer entities.CustomerRequest) error {
	return nil
}

func (csr *customerServiceImpl) GetCustomer(email string) (entities.CustomerRequest, error) {
	user, err := csr.repo.GetCustomerByEmail(email)
	if err != nil {
		return entities.CustomerRequest{}, err
	}

	userResp := entities.CustomerRequest{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userResp, nil
}
