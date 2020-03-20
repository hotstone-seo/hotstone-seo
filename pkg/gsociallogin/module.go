package gsociallogin

import (
	"github.com/typical-go/typical-go/pkg/typcore"
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
func (m *Module) Configure() *typcore.Configuration {
	return typcore.NewConfiguration(m.configName, &Config{})
}

// Provide the dependencies
func (m *Module) Provide() []*typdep.Constructor {
	return []*typdep.Constructor{
		typdep.NewConstructor(NewAuthGoogleService),
		typdep.NewConstructor(NewOauth2Config),
	}
}
