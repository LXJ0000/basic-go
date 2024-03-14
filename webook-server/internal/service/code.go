package service

import (
	"context"
	"math/rand"
	"strconv"
	"webook-server/internal/repository"
	"webook-server/internal/service/sms"
)

var ErrCodeSendFrequently = repository.ErrCodeSendFrequently

type codeService struct {
	repo repository.CodeRepository
	sms  sms.Service
}

func NewCodeService(repo repository.CodeRepository, sms sms.Service) CodeService {
	return &codeService{
		repo: repo,
		sms:  sms,
	}
}

// Send biz 区别使用业务 phone 接收者
func (svc *codeService) Send(ctx context.Context, biz string, phone string) error {
	code := svc.generateCode()
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}

	const codeTplId = "xxx"
	return svc.sms.Send(ctx, codeTplId, []string{code}, phone)
}

func (svc *codeService) Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	ok, err := svc.repo.Verify(ctx, biz, phone, inputCode)
	if err == repository.ErrCodeVerifyFrequently {
		return false, nil
	}
	return ok, err
}

func (svc *codeService) generateCode() string {
	num := rand.Intn(1000000)
	return strconv.Itoa(num)
}
