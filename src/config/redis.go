package cfg

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redOnce sync.Once
var client *redis.Client

func Red() *redis.Client {
	redOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     Cfg().Redis.Address,
			DB:       Cfg().Redis.Db,
			Password: Cfg().Secrets.RedisPassword,
		})
		if err := client.Ping(context.Background()).Err(); err != nil {
			panic("failed to connect to Redis: " + err.Error())
		}
	})
	return client
}
