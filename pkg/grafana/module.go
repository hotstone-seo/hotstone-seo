package grafana

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typdocker"
)

// Module of grafana
type Module struct {
	Port string
}

// New instance of grafana
func New() *Module {
	return &Module{
		Port: "3001",
	}
}

// WithPort to set port
func (m *Module) WithPort(port string) *Module {
	m.Port = port
	return m
}

// DockerCompose template
func (m *Module) DockerCompose(version typdocker.Version) *typdocker.ComposeObject {
	if version.IsV3() {
		return &typdocker.ComposeObject{
			Services: typdocker.Services{
				"grafana": typdocker.Service{
					Image: "grafana/grafana",
					Environment: map[string]string{
						"GF_SERVER_HTTP_PORT":         m.Port,
						"GF_SECURITY_ADMIN_PASSWORD":  "pass",
						"GF_SECURITY_ALLOW_EMBEDDING": "true",
						"GF_AUTH_ANONYMOUS_ENABLED":   "true",
					},
					// Volumes: []string{
					// 	"$PWD/grafana-data:/var/lib/grafana",
					// },
					Ports:    []string{fmt.Sprintf("%s:%s", m.Port, m.Port)},
					Networks: []string{"net-metrics"},
				},
			},
			Networks: typdocker.Networks{
				"net-metrics": typdocker.Network{
					Driver: "bridge",
				},
			},
		}
	}

	// TODO: create docker-compose for v2
	return nil
}
