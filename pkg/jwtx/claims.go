package jwtx

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
