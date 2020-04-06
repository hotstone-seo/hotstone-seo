package cachekit

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultMaxAge is default max-age value
	DefaultMaxAge = 30 * time.Second
)

// CacheControl is directive for cacche
type CacheControl struct {
	directives    []string
	defaultMaxAge time.Duration
}

// NewCacheControl return new instance of CacheControl
func NewCacheControl(directives ...string) *CacheControl {
	return &CacheControl{
		directives:    directives,
		defaultMaxAge: DefaultMaxAge,
	}
}

// CreateCacheControl to create new instance of CacheControl from request
func CreateCacheControl(req *http.Request) *CacheControl {
	var directives []string
	raw := req.Header.Get("Cache-Control")
	if raw != "" {
		for _, s := range strings.Split(raw, ",") {
			directives = append(directives, strings.TrimSpace(s))
		}
	}
	return NewCacheControl(directives...)
}

// WithDefaultMaxAge retunr CacheControl with new default max age
func (c *CacheControl) WithDefaultMaxAge(defaultMaxAge time.Duration) *CacheControl {
	c.defaultMaxAge = defaultMaxAge
	return c
}

// NoCache return true if no cache is set
func (c *CacheControl) NoCache() bool {
	for _, dir := range c.directives {
		if strings.ToLower(dir) == "no-cache" {
			return true
		}
	}
	return false
}

// MaxAge return max-age cache (in seconds)
func (c *CacheControl) MaxAge() time.Duration {
	for _, dir := range c.directives {
		dir = strings.ToLower(dir)
		keyName := "max-age="
		if strings.HasPrefix(dir, keyName) {
			maxAge, err := strconv.Atoi(dir[len(keyName):])
			if err != nil {
				break
			}
			return time.Duration(maxAge) * time.Second
		}
	}

	return c.defaultMaxAge
}

// Directives return directives for cache-control
func (c *CacheControl) Directives() []string {
	return c.directives
}
