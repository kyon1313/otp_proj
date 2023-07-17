package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < 6; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}
