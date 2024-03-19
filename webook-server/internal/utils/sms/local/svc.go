package local

import (
	"context"
	"log/slog"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) Send(ctx context.Context, templateId string, args []string, numbers ...string) error {
	slog.Info("Code is: ", args)
	return nil
}
