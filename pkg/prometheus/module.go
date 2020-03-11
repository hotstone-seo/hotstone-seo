package prometheus

import (
	"github.com/typical-go/typical-go/pkg/typdocker"
)

// Module of prometheus
type Module struct{}

// New instance of prometheus
func New() *Module {
	return &Module{}
}

// DockerCompose template
func (m *Module) DockerCompose(version typdocker.Version) *typdocker.ComposeObject {
	if version.IsV3() {
		return &typdocker.ComposeObject{
			Services: typdocker.Services{
				"prometheus": typdocker.Service{
					Image:   "prom/prometheus",
					Command: "--config.file=/etc/prometheus/prometheus.yml",
					Volumes: []string{
						"./metric/prometheus.yml:/etc/prometheus/prometheus.yml",
					},
					Ports:    []string{"9090:9090"},
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
