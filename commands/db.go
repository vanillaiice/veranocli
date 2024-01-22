package commands

import (
	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/veranocli"
)

// DB returns a command for handling database transactions in the CLI.
// It supports subcommands for inserting, deleting, getting, and updating activities in the database.
// The command takes optional flags for specifying the database path, file for insertion, and other options.
func DB() (cmd *cli.Command) {
	return &cli.Command{
		Name:    "db",
		Aliases: []string{"d"},
		Usage:   "make database transactions",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "db",
				Aliases: []string{"d"},
				Usage:   "open database `DB`",
				Value:   "activities.db",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:    "insert",
				Usage:   "insert activities in database",
				Aliases: []string{"i"},
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "parse and insert activities from `FILE`",
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "ignore",
						Aliases: []string{"i"},
						Usage:   "ignore duplicate inserts",
					},
					&cli.BoolFlag{
						Name:    "replace",
						Aliases: []string{"r"},
						Usage:   "replace duplicate inserts",
					},
				},
				Action: func(ctx *cli.Context) error {
					activities, err := veranocli.Parse(ctx.Path("file"))
					if err != nil {
						return err
					}

					if ctx.Bool("ignore") {
						err = veranocli.InsertActivities(activities, ctx.Path("db"), db.Ignore)
					} else if ctx.Bool("replace") {
						err = veranocli.InsertActivities(activities, ctx.Path("db"), db.Replace)
					} else {
						err = veranocli.InsertActivities(activities, ctx.Path("db"))
					}

					return err
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete activities in database",
				Flags: []cli.Flag{
					&cli.IntSliceFlag{
						Name:     "id",
						Usage:    "id of activities to delete",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					err := veranocli.DeleteActivities(ctx.Path("db"), ctx.IntSlice("id"))
					return err
				},
			},
			{
				Name:    "get",
				Usage:   "get activities in database",
				Aliases: []string{"g"},
				Flags: []cli.Flag{
					&cli.IntSliceFlag{
						Name:  "id",
						Usage: "ids of activities to get",
					},
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "get all activities in database",
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
						Usage:   "print raw activity structs",
					},
				},
				Action: func(ctx *cli.Context) error {
					var (
						activities []*activity.Activity
						err        error
					)

					if ctx.Bool("all") {
						activities, err = veranocli.GetActivitiesAll(ctx.Path("db"))
					} else {
						activities, err = veranocli.GetActivities(ctx.Path("db"), ctx.IntSlice("id"))
					}
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

					return nil
				},
			},
			{
				Name:    "update",
				Usage:   "update activity in database",
				Aliases: []string{"u"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "id",
						Usage: "id of the activity to update",
					},
					&cli.StringFlag{
						Name:    "field",
						Aliases: []string{"f"},
						Usage:   "activity field to update",
					},
					&cli.StringFlag{
						Name:    "value",
						Aliases: []string{"v"},
						Usage:   "new value of field",
					},
				},
				Action: func(ctx *cli.Context) error {
					err := veranocli.UpdateActivity(ctx.Path("db"), ctx.Int("id"), ctx.String("field"), ctx.String("value"))
					return err
				},
			},
		},
	}
}
