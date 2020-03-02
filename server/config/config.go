package config

// Config is configuration of app
type Config struct {
	Address                     string `default:":8089" required:"true"`
	CookieSecure                bool   `envconfig:"COOKIE_SECURE" default:"false"`
	JwtSecret                   string `envconfig:"JWT_SECRET"`
	Oauth2GoogleClientID        string `envconfig:"OAUTH2_GOOGLE_CLIENT_ID"`
	Oauth2GoogleClientSecret    string `envconfig:"OAUTH2_GOOGLE_CLIENT_SECRET"`
	Oauth2GoogleCallback        string `envconfig:"OAUTH2_GOOGLE_CALLBACK"`
	Oauth2GoogleRedirectSuccess string `envconfig:"OAUTH2_GOOGLE_REDIRECT_SUCCESS"`
	Oauth2GoogleRedirectFailure string `envconfig:"OAUTH2_GOOGLE_REDIRECT_FAILURE"`
	Oauth2GoogleHostedDomain    string `envconfig:"OAUTH2_GOOGLE_HOSTED_DOMAIN"`
	AuthLogoutRedirect          string `envconfig:"AUTH_LOGOUT_REDIRECT"`

	Debug bool `default:"false"`
}
