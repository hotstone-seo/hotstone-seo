package server

import (
	"github.com/hotstone-seo/hotstone-seo/server/config"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

// App of server
type App struct {
	configName string
}

// New application
func New() *App {
	return &App{
		configName: "APP",
	}
}

// WithConfigName return app module with new prefix
func (m *App) WithConfigName(configName string) *App {
	m.configName = configName
	return m
}

// Configure the application
func (m *App) Configure() *typcore.Configuration {
	return typcore.NewConfiguration(m.configName, &config.Config{})
}

// EntryPoint of application
func (*App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(main)
}

// Provide dependencies
func (m *App) Provide() []*typdep.Constructor {
	return []*typdep.Constructor{
		typdep.NewConstructor(typserver.New),
	}
}

// Destroy dependencies
func (m *App) Destroy() []*typdep.Invocation {
	return []*typdep.Invocation{
		typdep.NewInvocation(typserver.Shutdown),
	}
}

func main(s server, m taskManager) (err error) {
	if err = startTaskManager(m); err != nil {
		return
	}
	if err = startServer(s); err != nil {
		return
	}
	return
}
