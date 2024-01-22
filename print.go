package veranocli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/vanillaiice/verano/activity"
)

// PrintRaw prints the raw details of each activity in the provided slice.
func PrintRaw(activities []*activity.Activity) {
	for _, act := range activities {
		fmt.Printf("Activity %d: %+v\n", act.Id, act)
	}
}

// PrintPretty prints a formatted table with details of each activity in the provided slice.
// The 'withColor' parameter determines whether to use colored output.
func PrintPretty(activities []*activity.Activity, withColor bool) {
	tbl := table.New("Id", "Description", "Duration", "Start", "Finish", "Cost")
	if withColor {
		headerFormat := color.New(color.FgRed, color.Underline).SprintfFunc()
		columnFormat := color.New(color.FgYellow, color.Bold).SprintfFunc()
		tbl.WithHeaderFormatter(headerFormat).WithFirstColumnFormatter(columnFormat)
	}

	for _, act := range activities {
		tbl.AddRow(
			act.Id,
			act.Description,
			act.Duration.String(),
			act.Start.Format(TimeFormat),
			act.Finish.Format(TimeFormat),
			act.Cost,
		)
	}

	tbl.Print()
}
