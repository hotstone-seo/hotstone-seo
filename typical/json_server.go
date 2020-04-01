package typical

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

const (
	jsonServerSrc  = "scripts/json-server/db.json"
	jsonServerPort = "3021"
)

func jsonServer(c *typbuildtool.Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "json-server",
			Aliases: []string{"j"},
			Action: func(cliCtx *cli.Context) (err error) {

				ctx := cliCtx.Context
				jsonServer := "json-server"
				if !AvailableCommand(ctx, jsonServer) {
					c.Infof("Install %s", jsonServer)
					cmd := exec.CommandContext(ctx, "npm", "install", "-g", jsonServer)
					if err = cmd.Run(); err != nil {
						return
					}
				}

				c.Infof("Run %s", jsonServer)
				cmd := exec.CommandContext(ctx, jsonServer, "--port", jsonServerPort, jsonServerSrc)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				return cmd.Run()
			},
		},
	}
}

// AvailableCommand return path and boolean whether the command not available in bash
func AvailableCommand(ctx context.Context, name string) (ok bool) {
	if name == "" {
		return false
	}

	var debugger strings.Builder
	cmd := exec.CommandContext(ctx, "command", "-v", name)
	cmd.Stdout = &debugger

	if err := cmd.Run(); err != nil {
		return false
	}

	return strings.TrimSpace(debugger.String()) != ""
}
