package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	client *redis.Client
}

func NewRedisStorage(rdb *redis.Client) *Storage {
	return &Storage{client: rdb}
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *Storage) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return s.client.Set(ctx, key, value, ttl).Err()
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	return s.client.Del(ctx, key).Err()
}

func (s *Storage) Ping(ctx context.Context) error {
	return s.client.Ping(ctx).Err()
}
