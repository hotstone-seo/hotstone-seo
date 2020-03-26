package oauth2google

// Config is configuration of google social login
type Config struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	Callback     string `envconfig:"CALLBACK" required:"true"`
	HostedDomain string `envconfig:"HOSTED_DOMAIN"`

	CookieSecure bool   `envconfig:"COOKIE_SECURE" default:"false"`
	JWTSecret    string `envconfig:"JWT_SECRET"`

	RedirectSuccess string `envconfig:"REDIRECT_SUCCESS"`
	RedirectFailure string `envconfig:"REDIRECT_FAILURE"`
	LogoutRedirect  string `envconfig:"LOGOUT_REDIRECT"`
}
