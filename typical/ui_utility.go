package typical

import (
	"context"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

const (
	uiDir = "ui"
)

func uiUtility(*typbuildtool.Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "npm",
			Usage: "Execute `npm` command against ui source",
			Action: func(cli *cli.Context) error {
				return npm(cli.Context, cli.Args().Slice()...)
			},
		},
		{
			Name:  "ui",
			Usage: "UI Utility",
			Subcommands: []*cli.Command{
				{
					Name:    "start",
					Aliases: []string{"s"},
					Action: func(cli *cli.Context) error {
						return npm(cli.Context, "start")
					},
				},
				{
					Name:    "test",
					Aliases: []string{"t"},
					Action: func(cli *cli.Context) error {
						return npm(cli.Context, "run", "test:coverage")
					},
				},
				{
					Name:    "test-e2e",
					Aliases: []string{"e2e", "e"},
					Action: func(cli *cli.Context) error {
						return npm(cli.Context, "run", "test-e2e")
					},
				},
			},
		},
	}
}

func npm(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "npm", args...)
	cmd.Dir = uiDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
