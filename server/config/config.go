package config

// Config is configuration of app
type Config struct {
	Oauth2GoogleRedirectSuccess string `envconfig:"OAUTH2_GOOGLE_REDIRECT_SUCCESS"`
	Oauth2GoogleRedirectFailure string `envconfig:"OAUTH2_GOOGLE_REDIRECT_FAILURE"`
	Address            string `default:":8089" required:"true"`
	CookieSecure       bool   `envconfig:"COOKIE_SECURE" default:"false"`
	JwtSecret          string `envconfig:"JWT_SECRET"`
	AuthLogoutRedirect string `envconfig:"AUTH_LOGOUT_REDIRECT"`

	Debug bool `default:"false"`
}
