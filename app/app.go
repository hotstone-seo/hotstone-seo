package app

import (
	"github.com/hotstone-seo/hotstone-server/app/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli"
)

// Module of application
func Module() interface{} {
	return &module{
		Configuration: typcfg.Configuration{
			Prefix: "APP",
			Spec:   &config.Config{},
		},
	}
}

type module struct {
	typcfg.Configuration
}

func (module) AppCommands(c *typcli.ContextCli) []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:   "provider",
			Usage:  "Start the provider",
			Action: c.Action(startProvider),
		},
		cli.Command{
			Name:   "server",
			Usage:  "Start the Server",
			Action: c.Action(startServer),
		},
	}
}

func (m module) loadConfig(loader typcfg.Loader) (cfg config.Config, err error) {
	err = loader.Load(m.Configuration, &cfg)
	return
}
