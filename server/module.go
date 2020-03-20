package server

import (
	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

// Module of application
type Module struct {
	configName string
}

// New application
func New() *Module {
	return &Module{
		configName: "APP",
	}
}

// WithConfigName return app module with new prefix
func (m *Module) WithConfigName(configName string) *Module {
	m.configName = configName
	return m
}

// Configure the application
func (m *Module) Configure() *typcore.Configuration {
	return typcore.NewConfiguration(m.configName, &config.Config{})
}

// EntryPoint of application
func (*Module) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(func(s server, m taskManager) (err error) {
		if err = startTaskManager(m); err != nil {
			return
		}
		if err = startServer(s); err != nil {
			return
		}
		return
	})
}

// Provide dependencies
func (m *Module) Provide() []*typdep.Constructor {
	return []*typdep.Constructor{
		typdep.NewConstructor(typserver.New),
	}
}

// Destroy dependencies
func (m *Module) Destroy() []*typdep.Invocation {
	return []*typdep.Invocation{
		typdep.NewInvocation(typserver.Shutdown),
	}
}
