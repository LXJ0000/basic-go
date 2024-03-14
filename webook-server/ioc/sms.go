package ioc

import (
	"webook-server/internal/service/sms"
	"webook-server/internal/service/sms/local"
)

func InitSMSService() sms.Service {
	return local.NewService()
}
