package generators

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandonString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandonName() string {
	return RandonString(6)
}

func RandonEmail() string {
	var sb strings.Builder

	name := RandonString(6)
	domain := RandonString(6)

	sb.WriteString(name)
	sb.WriteRune('@')
	sb.WriteString(domain)
	sb.WriteString(".com")

	return sb.String()
}
