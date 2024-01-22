package commands

import (
	"veranocli"

	"github.com/goccy/go-graphviz"
	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/verano/graph"
	"github.com/vanillaiice/verano/util"
)

// Parse returns a command for parsing activities in various formats such as JSON, CSV, and XLSX.
// It supports options for specifying the file to parse, printing raw or pretty output, rendering a graph, and more.
func Parse() (cmd *cli.Command) {
	return &cli.Command{
		Name:    "parse",
		Aliases: []string{"p"},
		Usage:   "parse activities in json, csv, and xlsx formats",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "parse activities from `FILE`",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "do not print to stdout",
			},
			&cli.PathFlag{
				Name:    "graph",
				Aliases: []string{"g"},
				Usage:   "render graph of activities to `FILE`",
			},
			&cli.BoolFlag{
				Name:    "plain",
				Aliases: []string{"p"},
				Usage:   "print activities table with no colors",
			},
			&cli.BoolFlag{
				Name:    "raw",
				Aliases: []string{"r"},
				Usage:   "print raw activities",
			},
			&cli.PathFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "render image to `OUTPUT`",
			},
		},
		Action: func(ctx *cli.Context) error {
			activities, err := veranocli.Parse(ctx.Path("file"))
			if err != nil {
				return err
			}

			if !ctx.Bool("quiet") {
				if ctx.Bool("raw") {
					veranocli.PrintRaw(activities)
				} else {
					veranocli.PrintPretty(activities, !ctx.Bool("plain"))
				}
			}

			if ctx.Path("graph") != "" {
				g := graphviz.New()
				format, err := veranocli.GetFormat(ctx.Path("graph"))
				if err != nil {
					return err
				}
				activitiesMap := util.ActivitiesToMap(activities)
				err = graph.DrawAndRender(g, activitiesMap, format, ctx.Path("graph"))
			}

			return err
		},
	}
}
