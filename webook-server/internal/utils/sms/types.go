package sms

import (
	"context"
	"fmt"
	"math/rand"
)

type Service interface {
	Send(ctx context.Context, templateId string, args []string, numbers ...string) error
}

func GenerateCode() string {
	num := rand.Intn(1000000)
	return fmt.Sprintf("%06d", num)
}
