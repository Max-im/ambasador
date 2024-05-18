package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() error {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	return nil
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(3 * time.Second)
			Cache.Del(context.Background(), <-ch)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}

}
