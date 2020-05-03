package typical

import (
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/buildkit"
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
			Action:  c.ActionFunc("JSON-SERVER", startJSONServer),
		},
	}
}

func startJSONServer(c *typbuildtool.CliContext) (err error) {
	jsonServer := "json-server"
	if !buildkit.AvailableCommand(c, jsonServer) {
		c.Infof("Install %s", jsonServer)
		cmd := exec.CommandContext(c, "npm", "install", "-g", jsonServer)
		if err = cmd.Run(); err != nil {
			return
		}
	}

	c.Infof("Run %s", jsonServer)
	cmd := exec.CommandContext(c, jsonServer, "--port", jsonServerPort, jsonServerSrc)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
	return
}
