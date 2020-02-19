package config

// Config is configuration of app
type Config struct {
	Address                  string `default:":8089" required:"true"`
	MetricAddress            string `default:":8091"`
	DefaultLocale            string `default:"id_ID"`
	Oauth2GoogleClientID     string `envconfig:"OAUTH2_GOOGLE_CLIENT_ID"`
	Oauth2GoogleClientSecret string `envconfig:"OAUTH2_GOOGLE_CLIENT_SECRET"`
	Oauth2GoogleCallback     string `envconfig:"OAUTH2_GOOGLE_CALLBACK"`
}
