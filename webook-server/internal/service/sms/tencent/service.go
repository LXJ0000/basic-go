package tencent

import (
	"context"
	"fmt"
	"github.com/LXJ0000/go-lib/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	appId    *string
	signName *string
	client   *sms.Client
}

func NewService(appId string, signName string, client *sms.Client) *Service {
	return &Service{
		appId:    &appId,
		signName: &signName,
		client:   client,
	}
}

func (s Service) Send(ctx context.Context, templateId string, args []string, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = &templateId
	req.PhoneNumberSet = toStringPtrSlice(numbers)
	req.TemplateParamSet = toStringPtrSlice(args)

	resp, err := s.client.SendSms(req)
	if err != nil {

	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code != nil || *status.Code != "Ok" {
			return fmt.Errorf("Send SMS Error, Code: %s, Error: %s\n", *status.Code, *status.Message)
		}
	}
	return nil
}

func toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(src string) *string {
		return &src
	})
}
