package jsonreporter

import (
	"encoding/json"
	"fmt"

	"github.com/nikoheikkila/hours/report"
	"github.com/nikoheikkila/hours/toggl"
)

type JSONReport struct {
	entries []toggl.TimeEntry
}

func New(entries []toggl.TimeEntry) report.Exportable {
	return JSONReport{
		entries: entries,
	}
}

func (r JSONReport) Print() {
	document, err := json.Marshal(r.entries)

	if err != nil {
		fmt.Printf("Unable to convert entries to JSON. Error: '%s'\n", err)
	} else {
		fmt.Println(string(document))
	}
}
