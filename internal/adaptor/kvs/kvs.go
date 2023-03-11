package kvs

import (
	"context"
	"encoding/json"
	"errors"
	"gomulticache/internal/domain/cache"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/redis/go-redis/v9"
)

type KVSFactory[T any] interface {
	Of(*ristretto.Cache, *redis.Client, time.Duration) (*KVS[T], error)
}

var _ cache.Cache[any] = (*KVS[any])(nil)

type KVS[T any] struct {
	Cache  *ristretto.Cache
	Client *redis.Client
	TTL    time.Duration
}

func (kvs *KVS[T]) Get(ctx context.Context, key string) (T, error) {
	var value T

	// インメモリキャッシュから取得した値は[]byte型になっている

	val, ok := kvs.Cache.Get(key)
	if !ok {
		log.Printf("key not found: %s", key)
	}

	if _, ok := val.(T); ok {
		return val.(T), nil
	}

	// redisから取得した値は文字列になっている

	str, err := kvs.Client.Get(ctx, key).Result()
	if err != nil {
		return value, errors.New("1 failed to get cache")
	}

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		return value, errors.New("failed to unmarshal json")
	}

	return value, nil
}

func (kvs *KVS[T]) Set(ctx context.Context, key string, value T) error {
	ttl := time.Duration(time.Second)

	if ok := kvs.Cache.SetWithTTL(key, value, 1, ttl); !ok {
		return errors.New("failed to set in memory")
	}

	val, err := json.Marshal(value)
	if err != nil {
		return errors.New("failed to marshal json")
	}

	if err := kvs.Client.Set(ctx, key, string(val), ttl).Err(); err != nil {
		return errors.New("failed to set redis")
	}

	return nil
}

func (kvs *KVS[T]) Del(ctx context.Context, key string) error {
	kvs.Cache.Del(key)

	kvs.Client.Del(ctx, key)

	return nil
}
