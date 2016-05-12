package cli

import (
	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/fleet/create"
	"github.com/lcaballero/fleet/generate"
	"github.com/lcaballero/fleet/server"
)

func NewCli() *cmd.App {
	app := cmd.NewApp()
	app.Name = "fleet"
	app.Commands = []cmd.Command{
		serverCommand(),
		generateCommand(),
		newCommand(),
	}
	app.Action = generate.Generate
	return app
}

func newCommand() cmd.Command {
	return cmd.Command{
		Name:   "new",
		Action: create.New,
	}
}

func serverCommand() cmd.Command {
	return cmd.Command{
		Name:   "server",
		Action: server.Server,
	}
}

func generateCommand() cmd.Command {
	return cmd.Command{
		Name:   "generate",
		Usage:  "Runs the static site generation engine.",
		Action: generate.Generate,
	}
}
