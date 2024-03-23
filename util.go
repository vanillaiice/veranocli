package veranocli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-graphviz"
	"github.com/tealeg/xlsx/v3"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/verano/parser/pcsv"
	"github.com/vanillaiice/verano/parser/pjson"
	"github.com/vanillaiice/verano/parser/pxlsx"
)

// GetFormat determines the graphviz format based on the file extension.
// It returns the detected format and any encountered error during the process.
func GetFormat(file string) (format graphviz.Format, err error) {
	ext := filepath.Ext(file)
	switch ext {
	case ".dot":
		format = graphviz.XDOT
	case ".svg":
		format = graphviz.SVG
	case ".png":
		format = graphviz.PNG
	case ".jpg", ".jpeg":
		format = graphviz.JPG
	default:
		err = fmt.Errorf("extension %q not supported", ext)
	}
	return
}

// WriteFile writes the activities to a file with the specified 'filename'.
// The file format is determined based on the file extension.
// It returns any encountered error during the writing process.
func WriteFile(filename string, activities []*activity.Activity) (err error) {
	var file *os.File
	file, err = os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	ext := filepath.Ext(filename)
	switch ext {
	case ".db":
		err = InsertActivities(activities, filename, db.None)
	case ".json":
		err = pjson.ActivitiesToJSON(activities, file)
	case ".csv":
		err = pcsv.ActivitiesToCSV(activities, file)
	case ".xlsx":
		wb := xlsx.NewFile()
		sheet, err := wb.AddSheet("Sheet1")
		if err != nil {
			return err
		}
		pxlsx.ActivitiesToXLSX(activities, sheet)
		if err != nil {
			return err
		}
		wb.Save(filename)
	default:
		err = fmt.Errorf("file type %q not supported", ext)
	}

	return
}
