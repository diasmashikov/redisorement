package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx = context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
)
