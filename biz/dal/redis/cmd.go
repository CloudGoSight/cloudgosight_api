package redis

import (
	"context"
	"github.com/CloudGoSight/cloudgosight_api/util"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	util.Log().Info("redis set", key, value, expiration)
	if err := client.Set(ctx, key, value, expiration).Err(); err != nil {
		return errors.WithMessagef(err, "redis set %s %s", key, value)
	}
	return nil
}

func Get(ctx context.Context, key string) (string, bool, error) {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false, err
	} else if err != nil {
		return "", false, errors.WithMessagef(err, "redis get %s", key)
	}
	util.Log().Info("redis get", key, val)
	return val, true, nil
}
