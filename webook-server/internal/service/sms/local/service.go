package local

import (
	"context"
	"log"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) Send(ctx context.Context, templateId string, args []string, numbers ...string) error {
	log.Println("Code is: ", args)
	return nil
}
