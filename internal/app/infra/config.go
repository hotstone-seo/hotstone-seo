package infra

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	// postgres driver
	_ "github.com/lib/pq"
)

type (
	// App config
	App struct {
		Address string `default:":8089" required:"true"`
		Debug   bool   `default:"true"`
	}
	// Auth config
	Auth struct {
		JWTSecret       string `envconfig:"JWT_SECRET" require:"true"`
		ClientID        string `envconfig:"CLIENT_ID" required:"true"`
		ClientSecret    string `envconfig:"CLIENT_SECRET" required:"true"`
		Callback        string `envconfig:"CALLBACK" required:"true"`
		HostedDomain    string `envconfig:"HOSTED_DOMAIN"`
		CookieSecure    bool   `envconfig:"COOKIE_SECURE" default:"false"`
		RedirectSuccess string `envconfig:"REDIRECT_SUCCESS"`
		RedirectFailure string `envconfig:"REDIRECT_FAILURE"`
		LogoutRedirect  string `envconfig:"LOGOUT_REDIRECT" require:"true"`
	}

	// Redis Configuration
	Redis struct {
		Host     string `required:"true" default:"localhost"`
		Port     string `required:"true" default:"6379"`
		Password string `default:"redispass"`
		DB       int    `default:"0"`

		PoolSize           int           `envconfig:"POOL_SIZE"  default:"20" required:"true"`
		DialTimeout        time.Duration `envconfig:"DIAL_TIMEOUT" default:"5s" required:"true"`
		ReadWriteTimeout   time.Duration `envconfig:"READ_WRITE_TIMEOUT" default:"3s" required:"true"`
		IdleTimeout        time.Duration `envconfig:"IDLE_TIMEOUT" default:"5m" required:"true"`
		IdleCheckFrequency time.Duration `envconfig:"IDLE_CHECK_FREQUENCY" default:"1m" required:"true"`
		MaxConnAge         time.Duration `envconfig:"MAX_CONN_AGE" default:"30m" required:"true"`
	}
	// Pg is postgres configuration
	Pg struct {
		DBName   string `required:"true" default:"MyLibrary"`
		User     string `required:"true" default:"postgres"`
		Password string `required:"true" default:"pgpass"`
		Host     string `default:"localhost"`
		Port     string `default:"5432"`
	}
	// Analyt is timescale DB configuration
	Analyt struct {
		DBName   string `required:"true" default:"MyLibrary"`
		User     string `required:"true" default:"postgres"`
		Password string `required:"true" default:"pgpass"`
		Host     string `default:"localhost"`
		Port     string `default:"5432"`
	}
)

//
// Redis
//

func (r *Redis) connect() (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password:           r.Password,
		DB:                 r.DB,
		PoolSize:           r.PoolSize,
		DialTimeout:        r.DialTimeout,
		ReadTimeout:        r.ReadWriteTimeout,
		WriteTimeout:       r.ReadWriteTimeout,
		IdleTimeout:        r.IdleTimeout,
		IdleCheckFrequency: r.IdleCheckFrequency,
		MaxConnAge:         r.MaxConnAge,
	})

	if err = client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	return client, nil
}

//
// PG
//

func (p *Pg) connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DBName))
	if err != nil {
		return nil, fmt.Errorf("pg: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("pg: %w", err)
	}
	return db, nil
}

//
// Analyt
//

func (p *Analyt) connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DBName))
	if err != nil {
		return nil, fmt.Errorf("analyt: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("analyt: %w", err)
	}
	return db, nil
}
