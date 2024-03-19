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
)

type CodeCache interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type CodeCacheByRedis struct {
	cmd redis.Cmdable
}

func NewCodeCache(cmd redis.Cmdable) CodeCache {
	return &CodeCacheByRedis{cmd: cmd}
}

func (c *CodeCacheByRedis) Set(ctx context.Context, biz, phone, code string) error {
	res, err := c.cmd.Eval(ctx, luaSetCode, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}
	switch res {
	case -2:
		return err
	case -1:
		return ErrCodeSendFrequently
	}
	return nil
}

func (c *CodeCacheByRedis) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	res, err := c.cmd.Eval(ctx, luaVerifyCode, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return false, err
	}
	//ok 正确与否 err 是否还可以重试
	switch res {
	case -2: // 错误 还有机会
		return false, nil
	case -1: // 错误 没有机会
		return false, ErrCodeVerifyFrequently
	case 0: // 正确
		return true, nil
	}
	return false, err
}

func (c *CodeCacheByRedis) key(biz, phone string) string {
	return fmt.Sprintf("code:%s:%s", biz, phone) // code:login:[phone]
}
