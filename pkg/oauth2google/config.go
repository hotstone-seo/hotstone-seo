package oauth2google

import "github.com/typical-go/typical-go/pkg/typgo"

var (
	// DefaultConfigName is default config name of oauth2 google
	DefaultConfigName = "OAUTH2_GOOGLE"
)

// Config is configuration of google social login
type Config struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	Callback     string `envconfig:"CALLBACK" required:"true"`
	HostedDomain string `envconfig:"HOSTED_DOMAIN"`

	CookieSecure bool   `envconfig:"COOKIE_SECURE" default:"false"`

	RedirectSuccess string `envconfig:"REDIRECT_SUCCESS"`
	RedirectFailure string `envconfig:"REDIRECT_FAILURE"`
}

// Configuration of oauth2 google
func Configuration() *typgo.Configuration {
	return &typgo.Configuration{
		Name: DefaultConfigName,
		Spec: &Config{},
	}
}
