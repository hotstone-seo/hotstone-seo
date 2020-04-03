package cachekit

import (
	"net/http"
	"strings"
)

// CacheControl is directive for cacche
type CacheControl struct {
	directives []string
}

// NewCacheControl return new instance of CacheControl
func NewCacheControl(directives ...string) *CacheControl {
	return &CacheControl{
		directives: directives,
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

// NoCache return true if no cache is set
func (c *CacheControl) NoCache() bool {
	for _, dir := range c.directives {
		if strings.ToLower(dir) == "no-cache" {
			return true
		}
	}
	return false
}

// Directives return directives for cache-control
func (c *CacheControl) Directives() []string {
	return c.directives
}
