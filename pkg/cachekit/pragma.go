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

// Pragma handle pragmatic information/directives for caching
type Pragma struct {
	cacheControls []string
	defaultMaxAge time.Duration
}

// NewPragma return new instance of CacheControl
func NewPragma(cacheControls ...string) *Pragma {
	return &Pragma{
		cacheControls: cacheControls,
		defaultMaxAge: DefaultMaxAge,
	}
}

// CreatePragma to create new instance of CacheControl from request
func CreatePragma(req *http.Request) *Pragma {
	var directives []string
	raw := req.Header.Get("Cache-Control")
	if raw != "" {
		for _, s := range strings.Split(raw, ",") {
			directives = append(directives, strings.TrimSpace(s))
		}
	}
	return NewPragma(directives...)
}

// WithDefaultMaxAge retunr CacheControl with new default max age
func (c *Pragma) WithDefaultMaxAge(defaultMaxAge time.Duration) *Pragma {
	c.defaultMaxAge = defaultMaxAge
	return c
}

// NoCache return true if no cache is set
func (c *Pragma) NoCache() bool {
	for _, dir := range c.cacheControls {
		if strings.ToLower(dir) == "no-cache" {
			return true
		}
	}
	return false
}

// MaxAge return max-age cache (in seconds)
func (c *Pragma) MaxAge() time.Duration {
	for _, dir := range c.cacheControls {
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
func (c *Pragma) Directives() []string {
	return c.cacheControls
}
