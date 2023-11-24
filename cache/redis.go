package cache

import (
	"context"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func NewRedis() *Cache {
	return &Cache{redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
}

func (c *Cache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.Client.Get(ctx, key).Result()
}
