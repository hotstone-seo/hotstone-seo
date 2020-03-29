package oauth2google

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

var (
	// DefaultConfigName is default config name of oauth2 google
	DefaultConfigName = "OAUTH2_GOOGLE"
)

// Configuration the social login module
func Configuration() *typcfg.Configuration {
	return typcfg.NewConfiguration(DefaultConfigName, &Config{})
}

// Module of google oauth
func Module() *typapp.Module {
	return typapp.NewModule().
		WithProviders(
			typapp.NewConstructor(NewService),
		)
}
