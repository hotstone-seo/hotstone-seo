package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/dockerrx"
)

type pgDocker struct {
	name        string
	image       string
	userEnv     string
	passwordEnv string
	portEnv     string
}

var _ typdocker.Composer = (*pgDocker)(nil)

func (p *pgDocker) Compose() (*typdocker.Recipe, error) {
	pg := &dockerrx.Postgres{
		Version:  typdocker.V3,
		Name:     p.name,
		Image:    p.image,
		User:     os.Getenv(p.userEnv),
		Password: os.Getenv(p.passwordEnv),
		Port:     os.Getenv(p.portEnv),
	}
	return pg.Compose()
}
