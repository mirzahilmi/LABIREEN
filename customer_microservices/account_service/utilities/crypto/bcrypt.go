package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// DefaultCost factor used by bcrypt
const DefaultCost = 14

func HashValue(rawValue string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(rawValue), DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword := string(hashedPasswordBytes)
	return hashedPassword, nil
}

func CheckHash(rawValue, hashValue string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(rawValue))
	if err != nil {
		return err
	}

	return nil
}
