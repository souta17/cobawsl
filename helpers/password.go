package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func ComparePassword(hashing string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashing), []byte(password))
	return err
}
