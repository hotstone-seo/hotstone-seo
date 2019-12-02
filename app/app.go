package app

import (
	"github.com/hotstone-seo/hotstone-server/app/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli/v2"
)

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct{}

func (module) Commands(c *typcli.AppCli) []*cli.Command {
	return []*cli.Command{
		{Name: "provider", Usage: "Start the provider", Action: c.Action(startProvider)},
		{Name: "server", Usage: "Start the Server", Action: c.Action(startServer)},
	}
}

func (module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "APP"
	spec = &config.Config{}
	loadFn = func(loader typcfg.Loader) (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
