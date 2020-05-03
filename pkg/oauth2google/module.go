package oauth2google

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

var (
	// DefaultConfigName is default config name of oauth2 google
	DefaultConfigName = "OAUTH2_GOOGLE"
)

// Module of google oauth
func Module() *typapp.Module {
	return typapp.NewModule().
		Provide(typapp.NewConstructor("", NewService)).
		Configure(&typcfg.Configuration{
			Name: DefaultConfigName,
			Spec: &Config{},
		})
}
