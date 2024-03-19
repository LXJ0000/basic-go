package ioc

import (
	"webook-server/internal/utils/sms"
	"webook-server/internal/utils/sms/local"
)

func InitSMSService() sms.Service {
	return local.NewService()
}
