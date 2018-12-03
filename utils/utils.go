package utils

import (
	"math/rand"
	"strings"
	"strconv"
)

func RandomIP() string{
	parts := make([]string, 1, 4)
	parts[0] = "58"
	for i := 0; i < 3; i++ {
		parts = append(parts, strconv.Itoa(rand.Intn(256)))
	}
	return strings.Join(parts, ".")
}

func ValidateNumber(code string) bool{
	_, err := strconv.ParseInt(code, 10, 64)
	return err == nil && len(code) == 16
}