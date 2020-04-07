package cachekit

import (
	"fmt"
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
	noCache bool
	maxAge  time.Duration
	expires time.Time
}

// CreatePragma to create new instance of CacheControl from request
func CreatePragma(req *http.Request) *Pragma {
	var (
		noCache bool
	)

	maxAge := DefaultMaxAge

	if raw := req.Header.Get(HeaderCacheControl); raw != "" {
		for _, s := range strings.Split(raw, ",") {
			s = strings.ToLower(strings.TrimSpace(s))
			if s == "no-cache" {
				noCache = true
			}

			maxAgeField := "max-age="
			if strings.HasPrefix(s, maxAgeField) {
				maxAgeRaw, err := strconv.Atoi(s[len(maxAgeField):])
				if err != nil {
					break
				}
				maxAge = time.Duration(maxAgeRaw) * time.Second
			}
		}
	}
	return &Pragma{
		noCache: noCache,
		maxAge:  maxAge,
	}
}

// SetExpiresByTTL to set expires to current time to TTL
func (c *Pragma) SetExpiresByTTL(ttl time.Duration) {
	c.expires = time.Now().Add(ttl)
}

// NoCache return true if no cache is set
func (c *Pragma) NoCache() bool {
	return c.noCache
}

// MaxAge return max-age cache (in seconds)
func (c *Pragma) MaxAge() time.Duration {
	return c.maxAge
}

// ResponseHeaders return map that contain response header
func (c *Pragma) ResponseHeaders() map[string]string {
	var (
		m = make(map[string]string)
	)

	if !c.expires.IsZero() {
		m[HeaderExpires] = c.expires.Format(time.RFC1123)
	}

	if cacheControls := c.respCacheControls(); len(cacheControls) > 0 {
		m[HeaderCacheControl] = strings.Join(cacheControls, " ")
	}
	return m
}

func (c *Pragma) respCacheControls() (cc []string) {
	if c.NoCache() {
		cc = append(cc, "no-cache")
	} else {
		cc = append(cc, fmt.Sprintf("max-age=%d", int(c.MaxAge().Seconds())))
	}
	return
}
