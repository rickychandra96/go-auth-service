package util

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain text password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPassword compares a plain text password with a hashed password
func CheckPassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
