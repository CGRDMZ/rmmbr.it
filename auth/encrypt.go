package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFrom(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic("Something unexpected happened while hashing the password.")
	}

	return string(b)
}

func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
