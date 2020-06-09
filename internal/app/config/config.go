package config

// Config is configuration of app
type Config struct {
	Address        string `default:":8089" required:"true"`
	CookieSecure   bool   `envconfig:"COOKIE_SECURE" default:"false"`
	JWTSecret      string `envconfig:"JWT_SECRET"`
	LogoutRedirect string `envconfig:"LOGOUT_REDIRECT"`

	Debug bool `default:"false"`
}
