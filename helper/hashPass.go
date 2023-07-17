package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		log.Println("Unable to compare password", err)
		return false
	}
	return true
}
