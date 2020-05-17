package infra

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/typical-go/typical-rest-server/pkg/typpg"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"go.uber.org/dig"
)

type (
	// Configs of infra
	Configs struct {
		dig.In
		MainDB   *typpg.Config
		AnalytDB *typpg.Config `name:"analyt"`
		Redis    *typredis.Config
	}

	// Infras is list of infra to be provide in dependency-injection
	Infras struct {
		dig.Out
		MainDB   *sql.DB
		AnalytDB *sql.DB `name:"analyt"`
		Redis    *redis.Client
	}

	// Params infra
	Params struct {
		dig.In
		MainDB   *sql.DB
		AnalytDB *sql.DB `name:"analyt"`
		Redis    *redis.Client
	}
)

// Connect to infra
// @ctor
func Connect(c Configs) (infras Infras, err error) {
	var (
		mainDB   *sql.DB
		analytDB *sql.DB
		redis    *redis.Client
	)

	if mainDB, err = typpg.Connect(c.MainDB); err != nil {
		return
	}

	if analytDB, err = typpg.Connect(c.AnalytDB); err != nil {
		return
	}

	if redis, err = typredis.Connect(c.Redis); err != nil {
		return
	}

	return Infras{
		MainDB:   mainDB,
		AnalytDB: analytDB,
		Redis:    redis,
	}, nil
}

// Disconnect from postgres server
// @dtor
func Disconnect(p Params) (err error) {
	if err = typpg.Disconnect(p.MainDB); err != nil {
		return
	}
	if err = typpg.Disconnect(p.AnalytDB); err != nil {
		return
	}
	if err = typredis.Disconnect(p.Redis); err != nil {
		return
	}
	return
}
