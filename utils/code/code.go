package code

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateVerifyCode() (code string) {
	return strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899999) + 100000)
}
