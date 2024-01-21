package util

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashPassword, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
