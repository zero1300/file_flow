package load

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
		DB:   viper.GetInt("redis.db"),
	})
	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		panic("redis 连接失败....")
	}
	return client
}
