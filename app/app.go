package app

import (
	"github.com/hotstone-seo/hotstone-seo/app/config"
	"github.com/labstack/echo"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/typical-go/typical-rest-server/pkg/serverkit"
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
	return typdep.NewInvocation(func(s server, m TaskManager) error {
		if err := m.Start(); err != nil {
			return err
		}

		return s.Start()
	})
}

// Provide dependencies
func (m *Module) Provide() []*typdep.Constructor {
	return []*typdep.Constructor{
		typdep.NewConstructor(func(cfg config.Config) *echo.Echo {
			return serverkit.Create(cfg.Debug)
		}),
	}
}

// Destroy dependencies
func (m *Module) Destroy() []*typdep.Invocation {
	return []*typdep.Invocation{
		typdep.NewInvocation(serverkit.Shutdown),
	}
}
