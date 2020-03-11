package gsociallogin

// Config is configuration of google social login
type Config struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	Callback     string `envconfig:"CALLBACK" required:"true"`
	HostedDomain string `envconfig:"HOSTED_DOMAIN"`
}
