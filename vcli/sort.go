package vcli

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tealeg/xlsx/v3"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/verano/parser/pcsv"
	"github.com/vanillaiice/verano/parser/pjson"
	"github.com/vanillaiice/verano/parser/pxlsx"
	"github.com/vanillaiice/verano/project/timeline"
	"github.com/vanillaiice/verano/sorter"
	"github.com/vanillaiice/verano/util"
)

// Sort reads the file specified by 'filename' and organizes the activities based on their dependencies.
// The 'start' parameter represents the project start date.
// It returns a sorted slice of activities and any encountered error during the sorting process.
func Sort(filename, start string) (activitiesSorted []*activity.Activity, err error) {
	startDate, err := time.Parse(TimeFormat, start)
	if err != nil {
		return
	}

	var activities []*activity.Activity
	ext := filepath.Ext(filename)
	switch ext {
	case ".db":
		sqldb, err := db.New(filename)
		if err != nil {
			return nil, err
		}
		activities, err = sqldb.GetActivitiesAll()
	case ".json":
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		activities, err = pjson.JSONtoActivities(file)
	case ".csv":
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		activities, err = pcsv.CSVToActivities(file)
	case ".xlsx":
		wb, err := xlsx.OpenFile(filename)
		if err != nil {
			return nil, err
		}
		sheet := wb.Sheets[0]
		activities, err = pxlsx.XLSXToActivities(sheet)
	default:
		err = fmt.Errorf("file type %q not supported", ext)
	}

	if err != nil {
		return
	}

	activitiesGraph, err := util.ActivitiesToGraph(activities)
	if err != nil {
		return
	}
	order := sorter.SortActivitiesByDeps(activitiesGraph)
	activitiesMap := util.ActivitiesToMap(activities)
	activitiesSorted, err = sorter.SortActivitiesByOrder(activitiesMap, order)
	if err != nil {
		return
	}
	timeline.UpdateStartFinishTime(activitiesMap, order, startDate)

	return
}
