package jwtx

import (
	"labireen/customer_microservices/account_service/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomerClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func NewCustomerClaims(user entities.CustomerLogin, exp time.Duration) CustomerClaims {
	return CustomerClaims{
		Email:    user.Email,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
