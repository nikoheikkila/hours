package report

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nikoheikkila/hours/toggl"
)

func PrepareTable(entries []toggl.TimeEntry) table.Writer {
	markdown := table.NewWriter()
	markdown.SetOutputMirror(os.Stdout)
	markdown.AppendHeader(table.Row{"Date", "Task", "Project", "Client", "Hours"})

	for _, entry := range entries {
		markdown.AppendRow(table.Row{entry.FormatStartDate(), entry.Description, entry.GetProject(), entry.GetClient(), entry.FormatHours()})
	}

	return markdown
}
