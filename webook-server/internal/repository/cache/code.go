package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	//go:embed lua/set_code.lua
	luaSetCode string
	//go:embed lua/verify_code.lua
	luaVerifyCode string

	ErrCodeSendFrequently   = errors.New("验证码发送过于频繁")
	ErrCodeVerifyFrequently = errors.New("验证过于频繁")
	ErrUnknown              = errors.New("未知错误")
)

type CodeCache struct {
	cmd redis.Cmdable
}

func NewCodeCache(cmd redis.Cmdable) *CodeCache {
	return &CodeCache{cmd: cmd}
}

func (c *CodeCache) Set(ctx context.Context, biz, phone, code string) error {
	res, err := c.cmd.Eval(ctx, luaSetCode, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}
	switch res {
	case -2:
		return ErrUnknown
	case -1:
		return ErrCodeSendFrequently
	}
	return nil
}

func (c *CodeCache) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	res, err := c.cmd.Eval(ctx, luaVerifyCode, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return false, err
	}
	switch res {
	case -2:
		return false, err
	case -1:
		return false, ErrCodeVerifyFrequently
	case 0:
		return true, nil
	}
	return false, ErrUnknown
}

func (c *CodeCache) key(biz, phone string) string {
	return fmt.Sprintf("code:%s:%s", biz, phone) // code:login:[phone]
}
