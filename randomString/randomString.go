package randomstring

import (
	"math/rand"
	"time"
)

// Default charset
const charset_default = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Create(length int, charset string) string {
	if charset == "default" || charset == "" {
		charset = charset_default // Duplicate the default charset to argument charset
	}

	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = charset[seed.Intn(len(charset))]
	}

	return string(bytes)
}