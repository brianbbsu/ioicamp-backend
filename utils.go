package main

import (
	"crypto/rand"
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
