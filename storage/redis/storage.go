package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/hackfengJam/gokit/storage/api"
)

// Client implement api.Storage
// TODO tracing record
type Client struct {
	bizID string
	redis.Client
}

// NewStorage New Storage
func NewStorage(ctx context.Context, bizID string, cacheCli redis.Client) api.Storage {
	return &Client{
		bizID:  bizID,
		Client: cacheCli,
	}
}

// GetInstanceName Get InstanceName
func (c *Client) GetInstanceName() string {
	return c.bizID
}

// UnarySet Unary Set
func (c *Client) UnarySet(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	cmd := c.Set(key, value, expiration)
	return cmd.Err()
}

// MultiSet Multi Set
func (c *Client) MultiSet(ctx context.Context, keys []string, value map[string]interface{}, expiration time.Duration) (err error) {
	pipe := c.Pipeline()
	defer pipe.Close() // 释放

	// pipe.Set for loop
	for _, v := range keys {
		pipe.Set(v, value[v], expiration)
	}

	// 执行
	_, err = pipe.Exec()
	if err != nil {
		return
	}
	return nil
}

// Fetch fetch
func (c *Client) Fetch(ctx context.Context, key string) ([]byte, error) {
	cmd := c.Get(key)
	return cmd.Bytes()
}

// MultiFetch Multi Fetch
func (c *Client) MultiFetch(ctx context.Context, keys []string) ([]interface{}, error) {
	cmd := c.MGet(keys...)
	return cmd.Result()
}

// Delete delete
func (c *Client) Delete(ctx context.Context, keys []string) error {
	cmd := c.Del(keys...)
	return cmd.Err()
}
