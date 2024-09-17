package store

import (
	"math/rand/v2"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func genCode() string {
	const n = 8
	codeBytes := make([]byte, n)
	for i := range n {
		codeBytes[i] = characters[rand.IntN(len(characters))]
	}
	return string(codeBytes)
}
