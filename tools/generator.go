package tools

import (
	"math/rand"
	"strconv"
	"time"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func encodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		rem := num % 62
		result = string(base62Chars[rem]) + result
		num /= 62
	}
	return result
}

// GenerateShortID generates a unique, short, and URL-safe identifier.
// The function combines the current date, time (with seconds), nanoseconds, and a random value,
// then encodes it in Base62 to generate a compact and unique ID.
func GenerateShortID() string {
	now := time.Now()

	date := now.Format("02012006")
	time := now.Format("150405")
	nanoseconds := now.Nanosecond()

	combined := date + time + strconv.Itoa(nanoseconds)

	combinedInt, _ := strconv.ParseInt(combined[:15], 10, 64)

	randomValue := rand.Int63n(999999)
	finalValue := combinedInt + randomValue

	return encodeBase62(finalValue)
}