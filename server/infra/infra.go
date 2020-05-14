package infra

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"go.uber.org/dig"
)

type (
	// Configs of infra
	Configs struct {
		dig.In
		Pg    *typpostgres.Config
		Redis *typredis.Config
	}

	// Infras is list of infra to be provide in dependency-injection
	Infras struct {
		dig.Out
		Pg    *sql.DB
		Redis *redis.Client
	}

	// Params infra
	Params struct {
		dig.In
		Pg    *sql.DB
		Redis *redis.Client
	}
)

// Connect to infra
// @ctor
func Connect(c Configs) (infras Infras, err error) {
	var (
		pg    *sql.DB
		redis *redis.Client
	)

	if pg, err = typpostgres.Connect(c.Pg); err != nil {
		return
	}

	if redis, err = typredis.Connect(c.Redis); err != nil {
		return
	}

	return Infras{
		Pg:    pg,
		Redis: redis,
	}, nil
}

// Disconnect from postgres server
// @dtor
func Disconnect(p Params) (err error) {
	if err = typpostgres.Disconnect(p.Pg); err != nil {
		return
	}
	if err = typredis.Disconnect(p.Redis); err != nil {
		return
	}
	return
}
