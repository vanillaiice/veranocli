package vcli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/verano/util"
)

// Time format to use when parsing
const TimeFormat = "2 Jan 2006 15:04"

// GetActivities retrieves activities with the specified 'ids' from the database located at 'dbPath'.
// It returns a slice of activities and any encountered error during the operation.
func GetActivities(dbPath string, ids []int) (activities []*activity.Activity, err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}
	activities, err = sqldb.GetActivities(ids)
	return
}

// GetActivitiesAll retrieves all activities from the database located at 'dbPath'.
// It returns a slice of activities and any encountered error during the operation.
func GetActivitiesAll(dbPath string) (activities []*activity.Activity, err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}
	activities, err = sqldb.GetActivitiesAll()
	return
}

// InsertActivity inserts the provided 'act' activity into the database located at 'dbPath'.
// The optional 'duplicateInsertPolicy' parameter specifies the policy for handling duplicate inserts.
// It returns an error if the insertion operation encounters any issues.
func InsertActivity(act *activity.Activity, dbPath string, duplicateInsertPolicy ...db.DuplicateInsertPolicy) (err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}
	_, err = sqldb.InsertActivity(act, duplicateInsertPolicy...)
	return
}

// InsertActivities inserts the provided slice of activities 'acts' into the database located at 'dbPath'.
// The optional 'duplicateInsertPolicy' parameter specifies the policy for handling duplicate inserts.
// It returns an error if the insertion operation encounters any issues.
func InsertActivities(acts []*activity.Activity, dbPath string, duplicateInsertPolicy ...db.DuplicateInsertPolicy) (err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}
	err = sqldb.InsertActivities(acts, duplicateInsertPolicy...)
	return
}

// UpdateActivity updates the specified field of the activity with 'id' in the database located at 'dbPath'.
// The 'field' parameter specifies the field to be updated, and 'value' is the new value for that field.
// It returns an error if the update operation encounters any issues.
func UpdateActivity(dbPath string, id int, field string, value string) (err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}

	switch field {
	case "id":
		newId, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdateId(id, newId)
	case "description", "desc":
		_, err = sqldb.UpdateDescription(id, value)
	case "duration", "dur":
		newDuration, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdateDuration(id, newDuration)
	case "start", "st":
		newStart, err := time.Parse(TimeFormat, value)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdateStart(id, newStart)
	case "finish", "fi":
		newFinish, err := time.Parse(TimeFormat, value)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdateFinish(id, newFinish)
	case "predecessors", "pred":
		newPred, err := util.Unflat(value)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdatePredecessors(id, newPred)
	case "successors", "succ":
		newSucc, err := util.Unflat(value)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdateSuccessors(id, newSucc)
	case "cost":
		newCost, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		_, err = sqldb.UpdateCost(id, newCost)
	default:
		return fmt.Errorf("field %s not valid", field)
	}
	return err
}

// DeleteActivity deletes the activity with the specified 'id' from the database located at 'dbPath'.
// It returns an error if the deletion operation encounters any issues.
func DeleteActivity(dbPath string, id int) (err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}
	_, err = sqldb.DeleteActivity(id)
	return
}

// DeleteActivities deletes activities with the specified 'ids' from the database located at 'dbPath'.
// It returns an error if the deletion operation encounters any issues.
func DeleteActivities(dbPath string, id []int) (err error) {
	sqldb, err := db.New(dbPath)
	if err != nil {
		return
	}
	_, err = sqldb.DeleteActivities(id)
	return
}
