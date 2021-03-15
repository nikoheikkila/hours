package report

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nikoheikkila/hours/toggl"
	formats "github.com/nikoheikkila/hours/toggl/utils"
)

var headerRow = table.Row{"Date", "Task", "Project", "Client", "Hours"}

func PrepareTable(entries []toggl.TimeEntry) table.Writer {
	var totalMilliSeconds int64

	markdown := table.NewWriter()
	markdown.SetOutputMirror(os.Stdout)
	markdown.AppendHeader(headerRow)

	for _, entry := range entries {
		totalMilliSeconds += entry.Duration
		markdown.AppendRow(table.Row{entry.FormatStartDate(), entry.Description, entry.GetProject(), entry.GetClient(), entry.FormatHours()})
	}

	totalHours := formats.FormatDuration(formats.MillisecondsToHours(totalMilliSeconds))
	markdown.AppendFooter(table.Row{"Total", "", "", "", totalHours})

	return markdown
}
