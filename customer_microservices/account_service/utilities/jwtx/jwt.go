package jwtx

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(id uuid.UUID) (string, error) {
	tempExp := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(tempExp)

	if tempExp == "" || err != nil {
		exp = time.Hour * 1
	}

	tempToken := jwt.NewWithClaims(jwt.SigningMethodHS256, NewCustomerClaims(id, exp))

	token, err := tempToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeToken(signedToken string, ptrClaims jwt.Claims, secretKey string) error {

	token, err := jwt.ParseWithClaims(signedToken, ptrClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			// False signing method
			return nil, errors.New("wrong signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Parsing Error
		return errors.New("wrong signing method")
	}

	if !token.Valid {
		// Token not Valid
		return errors.New("unauthorized")
	}

	return nil
}
