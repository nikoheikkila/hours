package textreporter

import (
	"fmt"

	"github.com/muesli/termenv"
	"github.com/nikoheikkila/hours/report"
	"github.com/nikoheikkila/hours/toggl"
)

const DATE_LAYOUT string = "02.01.2006"

type PlainTextReport struct {
	entries []toggl.TimeEntry
}

func New(entries []toggl.TimeEntry) report.Exportable {
	return PlainTextReport{entries}
}

func (r PlainTextReport) Print() {
	p := termenv.ColorProfile()

	for _, entry := range r.entries {
		startDate := termenv.String(entry.FormatStartDate(DATE_LAYOUT)).Faint()
		description := termenv.String(entry.Description).Bold()
		project := termenv.String(entry.Project).Bold().Foreground(p.Color(entry.HexColor))
		hours := termenv.String(fmt.Sprintf("%.1f", entry.GetHours())).Bold()

		fmt.Printf("- %s: %s (%s) (%s h) \n", startDate, description, project, hours)
	}
}
