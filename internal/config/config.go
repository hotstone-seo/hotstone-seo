package config

// Config is configuration of app
type Config struct {
	Address string `default:":8089" required:"true"`

	Debug bool `default:"false"`
}
