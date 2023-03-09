package services

import (
	"errors"
	"labireen/customer_microservices/account_service/entities"
	"labireen/customer_microservices/account_service/repositories"
	"labireen/customer_microservices/account_service/utilities/crypto"

	"github.com/google/uuid"
)

type AuthService interface {
	RegisterCustomer(customer entities.CustomerRegister) error
	LoginCustomer(customer entities.CustomerLogin) (uuid.UUID, error)
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
		return errors.New("failed to encrypt given data")
	}

	assignID, err := uuid.NewRandom()
	if err != nil {
		return errors.New("failed to assign unique uuid")
	}

	user := entities.Customer{
		ID:               assignID,
		Name:             customer.Name,
		Email:            customer.Email,
		Password:         hashedPassword,
		PhoneNumber:      customer.PhoneNumber,
		VerificationCode: customer.VerificationCode,
	}

	err = asr.repo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (asr *authServiceImpl) LoginCustomer(customer entities.CustomerLogin) (uuid.UUID, error) {
	user, err := asr.repo.GetWhere("email", customer.Email)
	if err != nil {
		return uuid.UUID{}, errors.New("user not found")
	}

	if !user.Verified {
		return uuid.UUID{}, errors.New("user already verified")
	}

	if err := crypto.CheckHash(customer.Password, user.Password); err != nil {
		return uuid.UUID{}, errors.New("password is not valid or incorrect")
	}

	return user.ID, nil
}

func (asr *authServiceImpl) VerifyCustomer(code string) error {
	user, err := asr.repo.GetWhere("verification_code", code)
	if err != nil {
		return errors.New("user not found")
	}

	user.VerificationCode = ""
	user.Verified = true

	if err := asr.repo.Update(user); err != nil {
		return errors.New("failed to update user data")
	}

	return nil
}
