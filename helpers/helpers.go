package helpers

import (
	"math/rand"
)

func RandomNumber(n int) int {
	// rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
