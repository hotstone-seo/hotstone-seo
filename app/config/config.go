package config

// Config is configuration of app
type Config struct {
	Address       string `default:":8089" required:"true"`
	MetricAddress string `default:":8091"`
	DefaultLocale string `default:"id-ID"`
}
