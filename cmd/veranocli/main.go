package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/veranocli/commands"
)

func main() {
	app := &cli.App{
		Name:    "verano cli - manage activities in your project from the terminal",
		Suggest: true,
		Version: "v0.0.10",
		Authors: []*cli.Author{{Name: "Vanillaiice", Email: "vanillaiice1@proton.me"}},
		Usage:   "manage activities in a project",
		Commands: []*cli.Command{
			commands.Parse(),
			commands.DB(),
			commands.Sort(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
