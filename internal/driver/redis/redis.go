package redis

import (
	"context"
	"gomulticache/internal/adaptor/kvs"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/redis/go-redis/v9"
)

var _ kvs.KVSFactory[any] = (*Redis[any])(nil)

type Redis[T any] struct{}

func New[T any]() *Redis[T] {
	return &Redis[T]{}
}

func (rds *Redis[T]) Of(
	cache *ristretto.Cache,
	client *redis.Client,
	ttl time.Duration,
) (*kvs.KVS[T], error) {
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &kvs.KVS[T]{
		Cache:  cache,
		Client: client,
	}, nil
}

func NewClient(url string) *redis.Client {
	opt := &redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	return redis.NewClient(opt)
}

func NewCache() (*ristretto.Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}

	return cache, nil
}
