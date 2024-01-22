package vcli

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tealeg/xlsx/v3"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/parser/pcsv"
	"github.com/vanillaiice/verano/parser/pjson"
	"github.com/vanillaiice/verano/parser/pxlsx"
)

// Parse reads and interprets the content of a file specified by 'path' and returns a slice of activities.
// It automatically detects the file type based on its extension and uses the appropriate parser.
// Supported file types include SQLite database (.db), JSON (.json), CSV (.csv), and XLSX (.xlsx).
// It returns a slice of activities and any encountered error during the parsing process.
func Parse(path string) (activities []*activity.Activity, err error) {
	var reader *bytes.Reader
	ext := filepath.Ext(path)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}

	switch ext {
	case ".db":
		activities, err = GetActivitiesAll(path)
	case ".json":
		reader = bytes.NewReader(b)
		activities, err = pjson.JSONtoActivities(reader)
	case ".csv":
		reader = bytes.NewReader(b)
		activities, err = pcsv.CSVToActivities(reader)
	case ".xlsx":
		wb, err := xlsx.OpenBinary(b)
		if err != nil {
			return nil, err
		}
		sh := wb.Sheets[0]
		activities, err = pxlsx.XLSXToActivities(sh)
	default:
		err = fmt.Errorf("file type %q not supported", ext)
	}

	return
}
