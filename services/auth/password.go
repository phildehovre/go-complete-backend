package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	pw, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

func ComparePasswords(hashed string, submitted []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), submitted)
	return err
}
