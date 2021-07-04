package token

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandString(num int) string {
	res := make([]byte, num)
	for i := 0; i < num; i++ {
		res[i] = byte(RandInt(97, 122))
	}
	return string(res)
}
