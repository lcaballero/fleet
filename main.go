package main

import (
	"os"

	"github.com/lcaballero/fleet/cli"
)

func main() {
	cli.NewCli().Run(os.Args)
}
