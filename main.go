package main

import (
	"fmt"

	"github.com/muesli/termenv"
	"github.com/nikoheikkila/toggl-reporter/report"
)

func main() {
	var reports []report.Report

	reports = append(reports, report.New(1.2, "Test Project", "Meeting"))

	for i, report := range reports {
		index := termenv.String(fmt.Sprint(i + 1)).Bold().Foreground(termenv.ANSIBrightYellow)
		fmt.Printf("%s. %s", index, report.ToString())
	}
}