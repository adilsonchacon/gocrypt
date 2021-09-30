package gocrypt

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	var err error

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	} else {
		return string(hash), err
	}
}

func CheckPassword(password string, hashedPassword string) (bool, error) {
	var err error

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil, err
}

func RandString(size int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
