package cachekit

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
)

// Cache data
type Cache struct {
	key       string
	refreshFn RefreshFn
}

// RefreshFn is function that retrieve refresh data
type RefreshFn func() (interface{}, error)

// New cache
func New(key string, refreshFn RefreshFn) *Cache {
	return &Cache{
		key:       key,
		refreshFn: refreshFn,
	}
}

// Execute cache to retreive data and save to target variable
func (c *Cache) Execute(client *redis.Client, target interface{}, pragma *Pragma) error {
	var (
		v    interface{}
		data []byte
	)

	ttl, err := client.TTL(c.key).Result()
	if err != nil || ttl < 0 || pragma.NoCache() {

		if v, err = c.refreshFn(); err != nil {
			return fmt.Errorf("Cache: RefreshFunc: %w", err)
		}

		if data, err = c.marshal(v); err != nil {
			return fmt.Errorf("Cache: Marshal: %w", err)
		}

		ttl := pragma.MaxAge()

		if err = client.Set(c.key, data, ttl).Err(); err != nil {
			return fmt.Errorf("Cache: Set: %w", err)
		}

		pragma.SetExpiresByTTL(ttl)

		return copier.Copy(target, v)
	}

	pragma.SetExpiresByTTL(ttl)

	if data, err = client.Get(c.key).Bytes(); err != nil {
		return fmt.Errorf("Cache: Get: %w", err)
	}

	return c.unmarshal(data, target)
}

func (c *Cache) marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *Cache) unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
