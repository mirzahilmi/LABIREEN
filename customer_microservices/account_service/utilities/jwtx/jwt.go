package jwtx

import (
	"errors"
	"fmt"
	"labireen/customer_microservices/account_service/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user entities.CustomerLogin) (string, error) {
	tempExp := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(tempExp)

	if tempExp == "" || err != nil {
		exp = time.Hour * 1
	}

	tempToken := jwt.NewWithClaims(jwt.SigningMethodHS256, NewCustomerClaims(user, exp))

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
			return "", errors.New("wrong signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Parsing Error
		return fmt.Errorf("token has been tampered with")
	}

	if !token.Valid {
		// Token not Valid
		return fmt.Errorf("unauthorized")
	}

	return nil
}
