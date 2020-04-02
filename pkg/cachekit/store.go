package cachekit

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
)

// Store responsible to cache data
type Store struct {
	*redis.Client
	expiration time.Duration
}

// RefreshFn is function that handle refresh
type RefreshFn func() (interface{}, error)

// New return new instance of CacheStore
func New(client *redis.Client) *Store {
	return &Store{
		Client:     client,
		expiration: 100 * time.Millisecond,
	}
}

// WithExpiration return cache store with new expiration
func (c *Store) WithExpiration(expiration time.Duration) *Store {
	c.expiration = expiration
	return c
}

// Retrieve cache data
func (c *Store) Retrieve(ctx context.Context, key string, target interface{}, refresh RefreshFn) (fromCache bool, err error) {
	var data []byte
	if data, err = c.WithContext(ctx).Get(key).Bytes(); err != nil {
		var v interface{}
		if v, err = refresh(); err != nil {
			return
		}
		if data, err = c.marshal(v); err != nil {
			return
		}
		if err = c.WithContext(ctx).Set(key, data, c.expiration).Err(); err != nil {
			return
		}

		err = copier.Copy(target, v)
		return
	}

	return true, c.unmarshal(data, target)
}

func (c *Store) marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *Store) unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
