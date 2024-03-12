package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"webook-server/internal/domain"
)

var (
	ErrKeyNotExist = redis.Nil
)

type UserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:        client,
		expiration: time.Minute * 15,
	}
}

// Get 1. error = nil 则认为缓存击中 2. 如果没有数据返回特定 error 3. 如果系统出错 直接return error
func (c *UserCache) Get(ctx context.Context, userId int64) (domain.User, error) {
	key := c.Key(userId)
	val, err := c.cmd.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var user domain.User
	err = json.Unmarshal(val, &user)
	return user, err

}

func (c *UserCache) Set(ctx context.Context, user domain.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key := c.Key(user.UserId)
	return c.cmd.Set(ctx, key, val, c.expiration).Err()
}

func (c *UserCache) Key(userId int64) string {
	return fmt.Sprintf("user:info:%d", userId)
}
