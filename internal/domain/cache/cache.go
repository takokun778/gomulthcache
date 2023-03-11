package cache

import (
	"context"
)

type Cache[T any] interface {
	Get(context.Context, string) (T, error)
	Set(context.Context, string, T) error
	Del(context.Context, string) error
}
