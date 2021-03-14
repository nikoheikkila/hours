package markdownreporter

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nikoheikkila/hours/report"
	"github.com/nikoheikkila/hours/toggl"
)

type MarkdownReport struct {
	entries []toggl.TimeEntry
}

func New(entries []toggl.TimeEntry) report.Exportable {
	return MarkdownReport{
		entries: entries,
	}
}

func (r MarkdownReport) Print() {
	markdown := r.prepareTable()
	markdown.RenderMarkdown()
}

func (r MarkdownReport) prepareTable() table.Writer {
	markdown := table.NewWriter()
	markdown.SetOutputMirror(os.Stdout)
	markdown.AppendHeader(table.Row{"Date", "Task", "Project", "Hours"})

	for _, entry := range r.entries {
		markdown.AppendRow(table.Row{entry.FormatStartDate(), entry.Description, entry.Project, entry.FormatHours()})
	}

	return markdown
}
