package cachekit

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	// DefaultMaxAge is default max-age value
	DefaultMaxAge = 30
)

// CacheControl is directive for cacche
type CacheControl struct {
	directives    []string
	defaultMaxAge int // NOTE: in seconds
}

// NewCacheControl return new instance of CacheControl
func NewCacheControl(directives ...string) *CacheControl {
	return &CacheControl{
		directives:    directives,
		defaultMaxAge: DefaultMaxAge,
	}
}

// WithDefaultMaxAge retunr CacheControl with new default max age
func (c *CacheControl) WithDefaultMaxAge(defaultMaxAge int) *CacheControl {
	c.defaultMaxAge = defaultMaxAge
	return c
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
func (c *CacheControl) MaxAge() int {
	for _, dir := range c.directives {
		dir = strings.ToLower(dir)
		keyName := "max-age="
		if strings.HasPrefix(dir, keyName) {
			maxAge, err := strconv.Atoi(dir[len(keyName):])
			if err != nil {
				break
			}
			return maxAge
		}
	}

	return c.defaultMaxAge
}

// Directives return directives for cache-control
func (c *CacheControl) Directives() []string {
	return c.directives
}
