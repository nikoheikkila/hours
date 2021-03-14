package markdownreporter

import (
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
	report.PrepareTable(r.entries).RenderMarkdown()
}
