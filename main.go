package main

import (
	"fmt"
	"math"

	termenv "github.com/muesli/termenv"
)

type Report struct {
	hours float64
	project string
	description string
}

func NewReport(hours float64, project string, description string) Report {
	return Report{
		hours: math.Ceil(hours * 100) / 100,
		project: project,
		description: description,
	}
}

func (r Report) Hours() termenv.Style {
	out := fmt.Sprintf("%.1f", r.hours)
	return termenv.String(out).Bold().Foreground(termenv.ANSIMagenta)
}

func (r Report) Project() termenv.Style {
	return termenv.String(r.project).Bold().Foreground(termenv.ANSIBrightRed)
}

func (r Report) Description() termenv.Style {
	return termenv.String(r.description).Bold().Foreground(termenv.ANSIYellow)
}

func (r Report) ToString() termenv.Style {
	report := fmt.Sprintf("%s hours in %s with task '%s'", r.Hours(), r.Project(), r.Description())

	return termenv.String(report).Foreground(termenv.ANSIBrightGreen)
}

func main() {
	var reports []Report

	reports = append(reports, Report{1.20, "Test Project", "Meeting"})

	for i, report := range reports {
		index := termenv.String(fmt.Sprint(i + 1)).Bold().Foreground(termenv.ANSIBrightYellow)
		fmt.Printf("%s. %s", index, report.ToString())
	}
}