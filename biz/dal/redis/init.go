package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var (
	client *redis.Client
)

const TableName = "xxx"

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:        viper.GetString("redis.addr"),
		Password:    "", // 没有密码，默认值
		DB:          0,  // 默认DB 0
		ReadTimeout: time.Second * 3,
	})
}

func RedisKey(table string, key string) string {
	return fmt.Sprintf("[%s]%s", table, key)
}

func GenerateKey(key string) string {
	return RedisKey(TableName, key)
}
