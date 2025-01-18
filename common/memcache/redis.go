package memcache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient Client
)

type Client interface {
	Get(ctx context.Context, key string, out any) error
	Set(ctx context.Context, key string, in any) error
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key, value string, expriredDuration ...time.Duration) error
	Delete(ctx context.Context, key ...string) error
	MGetString(ctx context.Context, key ...string) ([]string, error)
	MSetString(ctx context.Context, keyVal map[string]string, keyExpireAt ...map[string]time.Time) error
}

type redisCache struct {
	*redis.Client
}

func (rc *redisCache) MGetString(ctx context.Context, keys ...string) ([]string, error) {
	ret := make([]string, 0)
	var errs = make([]error, 0)
	if cmd := rc.Client.MGet(ctx, keys...); cmd.Err() != nil {
		return nil, cmd.Err()
	} else {
		for idx, v := range cmd.Val() {
			if vStr, ok := v.(string); ok {
				ret = append(ret, vStr)
			} else {
				ret = append(ret, "")
				errs = append(errs, fmt.Errorf("value of %v is not string: %v", keys[idx], v))
				continue
			}
		}
	}

	return ret, errors.Join(errs...)
}

func (rc *redisCache) MSetString(ctx context.Context, keyval map[string]string, keyExpireAt ...map[string]time.Time) error {
	if err := rc.Client.MSet(ctx, keyval).Err(); err != nil {
		return err
	}

	if len(keyExpireAt) > 0 {
		var errs = make([]error, 0)
		for key, expireAt := range keyExpireAt[0] {
			if err := rc.Client.ExpireAt(ctx, key, expireAt).Err(); err != nil {
				errs = append(errs, fmt.Errorf("set EXPIREAT for %s: %s", key, err.Error()))
			}
		}
		if err := errors.Join(errs...); err != nil {
			return err
		}
	}

	return nil
}

func (rc *redisCache) GetString(ctx context.Context, key string) (string, error) {
	return rc.Client.Get(ctx, key).Result()
}

func (rc *redisCache) SetString(ctx context.Context, key, val string, expiredDuration ...time.Duration) error {
	if len(expiredDuration) == 0 {
		expiredDuration = []time.Duration{0}
	}
	return rc.Client.Set(ctx, key, val, expiredDuration[0]).Err()
}

func (rc *redisCache) Delete(ctx context.Context, key ...string) error {
	return rc.Client.Del(ctx, key...).Err()
}

func (rc *redisCache) Get(ctx context.Context, key string, out any) error {
	str, err := rc.GetString(ctx, key)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), out)
}

func (rc *redisCache) Set(ctx context.Context, key string, in any) error {
	bytes, err := json.Marshal(in)
	if err != nil {
		return err
	}
	if err := rc.SetString(ctx, key, string(bytes)); err != nil {
		return err
	}

	return nil
}

func GetClient() Client {
	return redisClient
}

func New(rdb *redis.Client) Client {
	return &redisCache{
		Client: rdb,
	}
}
