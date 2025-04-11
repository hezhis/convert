package convert

import (
	"math/rand"
)

// RandSlice 切片乱序
func RandSlice[T any](slice []T) {
	length := len(slice)
	if length < 2 {
		return
	}

	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return
}
