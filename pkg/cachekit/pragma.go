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

	// HeaderCacheControl is http header that holds directives (instructions) for caching in both requests and responses. A
	HeaderCacheControl = "Cache-Control"

	// HeaderExpires is http header that contains the date and time which denotes the period where the object can stay alive.
	HeaderExpires = "Expires"
)

// Pragma handle pragmatic information/directives for caching
type Pragma struct {
	cacheControls []string
	defaultMaxAge time.Duration
	expires       time.Time
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
	raw := req.Header.Get(HeaderCacheControl)
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

// SetExpiresByTTL to set expires to current time to TTL
func (c *Pragma) SetExpiresByTTL(ttl time.Duration) {
	c.expires = time.Now().Add(ttl)
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

// ResponseHeaders return map that contain response header
func (c *Pragma) ResponseHeaders() map[string]string {
	var (
		m = make(map[string]string)
	)

	if !c.expires.IsZero() {
		m[HeaderExpires] = c.expires.Format(time.RFC1123)
	}
	return m
}
