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
)

var (
	jsonServerHost = getEnv("JSON_SERVER_HOST", "localhost")
	jsonServerPort = getEnv("JSON_SERVER_PORT", "3021")
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
	cmd := exec.CommandContext(c, jsonServer, "--host", jsonServerHost, "--port", jsonServerPort, jsonServerSrc)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
	return
}

func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return defaultVal
}
