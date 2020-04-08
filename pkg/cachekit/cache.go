package cachekit

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
)

var (
	// ErrNoModified happen when conditional request apply
	ErrNoModified = errors.New("Cache: not modified")
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
func (c *Cache) Execute(client *redis.Client, target interface{}, pragma *Pragma) (err error) {
	var (
		v            interface{}
		modifiedTime time.Time
		ttl          time.Duration
	)

	if modifiedTime, err = c.getModifiedTime(client); err != nil {
		return fmt.Errorf("Cache: %w", err)
	}

	if ifModifiedTime := pragma.IfModifiedSince(); !ifModifiedTime.IsZero() && ifModifiedTime.After(modifiedTime) {
		err = ErrNoModified
		return
	}

	if modifiedTime.IsZero() || pragma.NoCache() {
		if v, err = c.refreshFn(); err != nil {
			return fmt.Errorf("Cache: RefreshFunc: %w", err)
		}

		ttl = pragma.MaxAge()

		if err = c.setData(client, v, ttl); err != nil {
			return fmt.Errorf("Cache: %w", err)
		}

		modifiedTime = time.Now()
		if err = c.setModifiedTime(client, modifiedTime, ttl); err != nil {
			return fmt.Errorf("Cache: %w", err)
		}

		pragma.SetLastModified(modifiedTime)
		pragma.SetExpiresByTTL(ttl)

		return copier.Copy(target, v)
	}

	if ttl, err = c.getData(client, target); err != nil {
		return fmt.Errorf("Cache: %w", err)
	}

	pragma.SetExpiresByTTL(ttl)
	pragma.SetLastModified(modifiedTime)

	return
}

func (c *Cache) setData(client *redis.Client, v interface{}, ttl time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(v); err != nil {
		return fmt.Errorf("Marshal: %w", err)
	}

	if err = client.Set(c.key, data, ttl).Err(); err != nil {
		return fmt.Errorf("Set: %w", err)
	}
	return
}

func (c *Cache) getData(client *redis.Client, target interface{}) (ttl time.Duration, err error) {
	var data []byte

	if ttl, err = client.TTL(c.key).Result(); err != nil {
		err = fmt.Errorf("TTL: %w", err)
		return
	}

	if data, err = client.Get(c.key).Bytes(); err != nil {
		err = fmt.Errorf("Get: %w", err)
		return
	}

	if err = json.Unmarshal(data, target); err != nil {
		err = fmt.Errorf("Unmarsal: %w", err)
		return
	}

	return
}

func (c *Cache) setModifiedTime(client *redis.Client, t time.Time, ttl time.Duration) (err error) {
	if err = client.Set(c.modifiedTimeKey(), t.Format(time.RFC1123), ttl).Err(); err != nil {
		return fmt.Errorf("SetModifiedTime: %w", err)
	}
	return
}

func (c *Cache) getModifiedTime(client *redis.Client) (modified time.Time, err error) {
	var (
		raw string
	)

	if raw = client.Get(c.modifiedTimeKey()).Val(); raw == "" {
		return
	}

	if modified, err = time.Parse(time.RFC1123, raw); err != nil {
		err = fmt.Errorf("ParseModifiedTime: %w", err)
		return
	}

	return
}

func (c *Cache) modifiedTimeKey() string {
	return c.key + ":time"
}
