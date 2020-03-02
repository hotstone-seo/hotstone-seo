package server

import (
	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

// Module of application
type Module struct {
	prefix string
}

// New application
func New() *Module {
	return &Module{
		prefix: "APP",
	}
}

// WithPrefix return app module with new prefix
func (m *Module) WithPrefix(prefix string) *Module {
	m.prefix = prefix
	return m
}

// Configure the application
func (m *Module) Configure(loader typcfg.Loader) *typcfg.Detail {
	return &typcfg.Detail{
		Prefix: m.prefix,
		Spec:   &config.Config{},
		Constructor: typdep.NewConstructor(func() (cfg config.Config, err error) {
			err = loader.Load(m.prefix, &cfg)
			return
		}),
	}
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
