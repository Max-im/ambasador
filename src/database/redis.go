package database

import "github.com/redis/go-redis/v9"

var Cache *redis.Client

func SetupRedis() error {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	return nil
}
