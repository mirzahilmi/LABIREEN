package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomerClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func NewCustomerClaims(id uuid.UUID, exp time.Duration) CustomerClaims {
	return CustomerClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
