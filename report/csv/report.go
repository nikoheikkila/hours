package csvreporter

import (
	"github.com/nikoheikkila/hours/report"
	"github.com/nikoheikkila/hours/toggl"
)

type CSVReport struct {
	entries []toggl.TimeEntry
}

func New(entries []toggl.TimeEntry) report.Exportable {
	return CSVReport{
		entries: entries,
	}
}

func (r CSVReport) Print() {
	report.PrepareTable(r.entries).RenderCSV()
}
