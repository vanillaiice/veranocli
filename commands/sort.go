package commands

import (
	"time"
	"veranocli"

	"github.com/goccy/go-graphviz"
	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/verano/graph"
	"github.com/vanillaiice/verano/util"
)

// Sort returns a command for topologically sorting activities.
// It supports options for specifying the file to parse, the start date of the project, printing raw or pretty output,
// outputting sorted activities to a file, and rendering a graph.
func Sort() (cmd *cli.Command) {
	return &cli.Command{
		Name:    "sort",
		Aliases: []string{"s"},
		Usage:   "topologically sort activities",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "parse activities from `FILE`",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "start",
				Aliases: []string{"d"},
				Usage:   "start date of the project",
				Value:   time.Now().Format(veranocli.TimeFormat),
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "do not print to stdout",
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
				Usage:   "output sorted activities to `FILE`",
			},
			&cli.PathFlag{
				Name:    "graph",
				Aliases: []string{"g"},
				Usage:   "render graph of activities to `FILE`",
			},
		},
		Action: func(ctx *cli.Context) error {
			activities, err := veranocli.Sort(ctx.Path("file"), ctx.String("start"))
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

			if ctx.Path("output") != "" {
				err = veranocli.WriteFile(ctx.Path("output"), activities)
			}

			if ctx.Path("graph") != "" {
				format, err := veranocli.GetFormat(ctx.Path("graph"))
				if err != nil {
					return err
				}
				activitiesMap := util.ActivitiesToMap(activities)
				g := graphviz.New()
				err = graph.DrawAndRender(g, activitiesMap, format, ctx.Path("graph"))
			}

			return err
		},
	}
}
