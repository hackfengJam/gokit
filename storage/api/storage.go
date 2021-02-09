package api

import (
	"context"
	"time"
)

// Storage interface
type Storage interface {
	GetInstanceName() string
	UnarySet(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	MultiSet(ctx context.Context, keys []string, value map[string]interface{}, expiration time.Duration) error
	Fetch(ctx context.Context, key string) ([]byte, error)
	MultiFetch(ctx context.Context, keys []string) ([]interface{}, error)
	Delete(ctx context.Context, keys []string) error
}
