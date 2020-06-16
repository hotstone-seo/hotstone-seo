package infra

import (
	"database/sql"

	"github.com/go-redis/redis"
	"go.uber.org/dig"
)

type (
	// Configs of infra
	Configs struct {
		dig.In
		MainDB *Pg
		Analyt *Analyt
		Redis  *Redis
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

	if mainDB, err = c.MainDB.connect(); err != nil {
		return
	}

	if analytDB, err = c.Analyt.connect(); err != nil {
		return
	}

	if redis, err = c.Redis.connect(); err != nil {
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
	if err = p.MainDB.Close(); err != nil {
		return
	}
	if err = p.AnalytDB.Close(); err != nil {
		return
	}
	if err = p.Redis.Close(); err != nil {
		return
	}
	return
}
