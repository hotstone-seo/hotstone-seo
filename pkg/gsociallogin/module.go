package gsociallogin

import (
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typdep"
)

// Module for Social Login
type Module struct {
	configName string
}

// New instance of module
func New() *Module {
	return &Module{
		configName: "OAUTH2_GOOGLE",
	}
}

// Configure the social login module
func (m *Module) Configure(loader typcfg.Loader) *typcfg.Configuration {
	return &typcfg.Configuration{
		Name: m.configName,
		Spec: &Config{},
		Constructor: typdep.NewConstructor(func() (cfg Config, err error) {
			err = loader.Load(m.configName, &cfg)
			return
		}),
	}
}

// Provide the dependencies
func (m *Module) Provide() []*typdep.Constructor {
	return []*typdep.Constructor{
		typdep.NewConstructor(NewAuthGoogleService),
		typdep.NewConstructor(NewOauth2Config),
	}
}
