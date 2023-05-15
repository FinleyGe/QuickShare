package redis

import (
	"context"
)

func SetHash(shareCode string, hash string) error {
	ctx := context.Background()
	return Redis.Set(ctx, shareCode, hash, 0).Err()
}

func GetHash(shareCode string) (string, error) {
	ctx := context.Background()
	return Redis.Get(ctx, shareCode).Result()
}
