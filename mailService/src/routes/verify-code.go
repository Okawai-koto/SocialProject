package routes

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func verifycode() string {

	verify := make([]byte, 5)

	for i := range verify {
		verify[i] = charset[rand.Intn(len(charset))]
	}

	return string(verify)
}
