package oauth2google

import "time"

var (
	// CookieExpiration is expiration for JWT cookie
	CookieExpiration time.Duration = 72 * time.Hour

	StateExpiration time.Duration = 20 * time.Minute

	TokenEpiration time.Duration = 72 * time.Hour
)
