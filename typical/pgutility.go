package typical

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"

	"github.com/typical-go/typical-rest-server/pkg/pgutil"
)

type pgUtility struct {
	name         string
	hostEnv      string
	portEnv      string
	userEnv      string
	passwordEnv  string
	dbnameEnv    string
	migrationSrc string
	seedSrc      string
}

var _ typgo.Utility = (*pgUtility)(nil)

func (u *pgUtility) Commands(c *typgo.BuildCli) ([]*cli.Command, error) {
	return []*cli.Command{
		{
			Name:  u.name,
			Usage: fmt.Sprintf("%s utility", u.name),
			Subcommands: []*cli.Command{
				{
					Name:  "create",
					Usage: "Create database",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						return u.execute(c, "create")
					}),
				},
				{
					Name:  "drop",
					Usage: "Drop database",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						return u.execute(c, "drop")
					}),
				},
				{
					Name:  "migrate",
					Usage: "Migrate database",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						return u.execute(c, "migrate")
					}),
				},
				{
					Name:  "rollback",
					Usage: "Rollback database",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						return u.execute(c, "rollback")
					}),
				},
				{
					Name:  "seed",
					Usage: "Seed database",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						return u.execute(c, "rollback")
					}),
				},
				{
					Name:  "reset",
					Usage: "Reset database",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						if err := u.execute(c, "drop"); err != nil {
							return err
						}
						if err := u.execute(c, "create"); err != nil {
							return err
						}
						if err := u.execute(c, "migrate"); err != nil {
							return err
						}
						if err := u.execute(c, "seed"); err != nil {
							return err
						}
						return nil
					}),
				},
				{
					Name:  "console",
					Usage: "Postgres console",
					Action: c.ActionFn("PG", func(c *typgo.Context) error {
						return c.Execute(&pgutil.PgConsole{
							Host:     os.Getenv(u.hostEnv),
							Port:     os.Getenv(u.portEnv),
							User:     os.Getenv(u.userEnv),
							Password: os.Getenv(u.passwordEnv),
						})
					}),
				},
			},
		},
	}, nil
}

func (u *pgUtility) execute(c *typgo.Context, action string) error {
	bin := ".typical-tmp/bin/pgutil"
	if _, err := os.Stat(bin); os.IsNotExist(err) {
		if err := c.Execute(&buildkit.GoBuild{
			Out:    bin,
			Source: "github.com/typical-go/typical-rest-server/pkg/pgutil/cli",
		}); err != nil {
			return err
		}
	}

	return c.Execute(&pgutil.PgUtil{
		Name:         bin,
		Action:       action,
		Host:         os.Getenv(u.hostEnv),
		Port:         os.Getenv(u.portEnv),
		User:         os.Getenv(u.userEnv),
		Password:     os.Getenv(u.passwordEnv),
		DBName:       os.Getenv(u.dbnameEnv),
		MigrationSrc: u.migrationSrc,
		SeedSrc:      u.seedSrc,
	})
}
