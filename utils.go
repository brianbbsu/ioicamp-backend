package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func getRandomToken(n int) (string, error) {
	const sigma = "23456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	ret := make([]byte, 0)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(sigma))))
		if err != nil {
			return "", err
		}
		ret = append(ret, sigma[num.Int64()])
	}
	return string(ret), nil
}

func validateNewPassword(password string) error {
	// Check if password is strong enough
	// TODO: Implement this
	if len(password) == 0 {
		return errors.New("Password should have length at least 1")
	} else if len([]byte(password)) > 72 { // 72 is the maximnum supported length for bcrypt
		return errors.New("Password should have length not greater than 72")
	}
	return nil
}
