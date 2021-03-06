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
			startDate := termenv.String(entry.FormatStartDate()).Faint()
			description := termenv.String(entry.GetDescription()).Bold()
			project := termenv.String(entry.GetProject()).Bold().Foreground(profile.Color(entry.HexColor))
			client := termenv.String(entry.GetClient()).Bold().Foreground(profile.Color(entry.HexColor))
			hours := termenv.String(fmt.Sprintf("%.1f", entry.GetHours())).Bold()

			fmt.Printf("%s: %s (%s, %s) (%s h) \n", startDate, description, project, client, hours)
		} else {
			startDate := entry.FormatStartDate()
			description := entry.GetDescription()
			project := entry.GetProject()
			client := entry.GetClient()
			hours := fmt.Sprintf("%.1f", entry.GetHours())

			fmt.Printf("%s: %s (%s, %s) (%s h) \n", startDate, description, project, client, hours)
		}
	}
}
