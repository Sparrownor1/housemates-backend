package group

import (
	"math/rand"
	"time"
)

const inviteCodeLength = 6

var numbers = []rune("0123456789")

func GenerateInviteCode() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, inviteCodeLength)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}

	return string(b)
}
