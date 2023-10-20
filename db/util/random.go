// Generation of random data
package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqurstuvxyz"

func init() {
	// Doubt
	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
}

// Not used, just additional info
// Generates random int b/w min and max
func RandomInt(min, max int64) int64 {
	// Int 63n generates b/w 0 - arg, here, we add min, so it will be from min + max - min + 1 -> min - max
	return min + rand.Int63n(max-min+1)
}

// Generates random string of len n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "@chat.io"
}
