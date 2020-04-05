package cachekit

import (
	"encoding/json"

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
func (c *Cache) Execute(client *redis.Client, target interface{}, cc *CacheControl) error {
	data, err := client.Get(c.key).Bytes()
	if err != nil || len(data) <= 0 || cc.NoCache() {
		var v interface{}
		if v, err = c.refreshFn(); err != nil {
			return err
		}
		if data, err = c.marshal(v); err != nil {
			return err
		}
		if err = client.Set(c.key, data, cc.MaxAge()).Err(); err != nil {
			return err
		}
		return copier.Copy(target, v)
	}

	return c.unmarshal(data, target)
}

func (c *Cache) marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *Cache) unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
