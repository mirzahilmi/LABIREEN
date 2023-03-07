package services

import (
	"fmt"
	"labireen/customer_microservices/account_service/entities"
	"labireen/customer_microservices/account_service/repositories"
	"labireen/customer_microservices/account_service/utilities/crypto"

	"github.com/google/uuid"
)

type AuthService interface {
	RegisterCustomer(customer entities.CustomerRegister) error
	LoginCustomer(customer entities.CustomerLogin) error
	VerifyCustomer(email string) error
}

type authServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authServiceImpl{repo}
}

func (asr *authServiceImpl) RegisterCustomer(customer entities.CustomerRegister) error {
	hashedPassword, err := crypto.HashValue(customer.Password)
	if err != nil {
		return err
	}

	assignID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user := entities.Customer{
		ID:               assignID,
		Name:             customer.Name,
		Email:            customer.Email,
		Password:         hashedPassword,
		PhoneNumber:      customer.PhoneNumber,
		VerificationCode: customer.VerificationCode,
	}

	err = asr.repo.CreateCustomer(&user)
	if err != nil {
		return err
	}

	return nil
}

func (asr *authServiceImpl) LoginCustomer(customer entities.CustomerLogin) error {
	user, err := asr.repo.GetCustomerByEmail(customer.Email)
	if err != nil {
		return err
	}

	if !user.Verified {
		return fmt.Errorf("please verify your email")
	}

	if err := crypto.CheckHash(customer.Password, user.Password); err != nil {
		return err
	}

	return nil
}

func (asr *authServiceImpl) VerifyCustomer(code string) error {
	user, err := asr.repo.GetCustomerByCustom("verification_code", code)
	if err != nil {
		return err
	}

	user.VerificationCode = ""
	user.Verified = true
	if err := asr.repo.UpdateCustomer(user); err != nil {
		return err
	}

	return nil
}
