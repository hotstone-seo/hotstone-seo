package typical

import (
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

const (
	jsonServerSrc = "scripts/json-server/db.json"
)

var (
	jsonServerHost = getEnv("JSON_SERVER_HOST", "localhost")
	jsonServerPort = getEnv("JSON_SERVER_PORT", "3021")
)

func jsonServer(b *typgo.BuildTool) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "json-server",
			Aliases: []string{"j"},
			Action:  b.ActionFunc("JSON-SERVER", startJSONServer),
		},
	}
}

func startJSONServer(c *typgo.Context) (err error) {
	jsonServer := "json-server"
	ctx := c.Cli.Context
	if !buildkit.AvailableCommand(ctx, jsonServer) {
		c.Infof("Install %s", jsonServer)
		cmd := exec.CommandContext(ctx, "npm", "install", "-g", jsonServer)
		if err = cmd.Run(); err != nil {
			return
		}
	}

	c.Infof("Run %s", jsonServer)
	cmd := exec.CommandContext(ctx, jsonServer, "--host", jsonServerHost, "--port", jsonServerPort, jsonServerSrc)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
