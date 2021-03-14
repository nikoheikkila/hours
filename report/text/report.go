package textreporter

import (
	"fmt"

	"github.com/muesli/termenv"
	"github.com/nikoheikkila/hours/report"
	"github.com/nikoheikkila/hours/toggl"
)

const DATE_LAYOUT string = "02.01.2006"

type PlainTextReport struct {
	entries   []toggl.TimeEntry
	colorized bool
}

func NewColorized(entries []toggl.TimeEntry) report.Exportable {
	return PlainTextReport{
		entries:   entries,
		colorized: true,
	}
}

func New(entries []toggl.TimeEntry) report.Exportable {
	return PlainTextReport{
		entries:   entries,
		colorized: false,
	}
}

func (r PlainTextReport) Print() {
	for _, entry := range r.entries {
		if r.colorized {
			profile := termenv.ColorProfile()
			startDate := termenv.String(entry.FormatStartDate(DATE_LAYOUT)).Faint()
			description := termenv.String(entry.Description).Bold()
			project := termenv.String(entry.Project).Bold().Foreground(profile.Color(entry.HexColor))
			hours := termenv.String(fmt.Sprintf("%.1f", entry.GetHours())).Bold()

			fmt.Printf("- %s: %s (%s) (%s h) \n", startDate, description, project, hours)
		} else {
			startDate := entry.FormatStartDate(DATE_LAYOUT)
			description := entry.Description
			project := entry.Project
			hours := fmt.Sprintf("%.1f", entry.GetHours())

			fmt.Printf("- %s: %s (%s) (%s h) \n", startDate, description, project, hours)
		}
	}
}
