package app

import (
	"github.com/hotstone-seo/hotstone-seo/app/config"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// New application [nowire]
func New() *Module {
	return &Module{}
}

// Module of application
type Module struct{}

// EntryPoint of application
func (*Module) EntryPoint() interface{} {
	return func(s server,  m TaskManager) error {
		if err := m.Start(); err != nil {
			return err
		}


		s.Middleware()
		s.Route()
		return s.Start()
	}
}

// Configure the application
func (*Module) Configure(loader typcore.ConfigLoader) (prefix string, spec, loadFn interface{}) {
	prefix = "APP"
	spec = &config.Config{}
	loadFn = func() (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
