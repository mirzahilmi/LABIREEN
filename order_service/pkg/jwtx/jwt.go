package jwtx

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

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
