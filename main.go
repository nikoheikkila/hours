package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/nikoheikkila/hours/report"
	csvreporter "github.com/nikoheikkila/hours/report/csv"
	jsonreporter "github.com/nikoheikkila/hours/report/json"
	markdownreporter "github.com/nikoheikkila/hours/report/markdown"
	textreporter "github.com/nikoheikkila/hours/report/text"
	"github.com/nikoheikkila/hours/rules"
	"github.com/nikoheikkila/hours/toggl"
)

const DATE_FORMAT_ISO string = "2006-01-02"

func main() {
	var err error

	configuration, err := toggl.LoadConfiguration()
	handleError(err)

	client := toggl.New(configuration)

	output := flag.String("output", "text", "Output format for the reporter.")
	since := flag.String("since", time.Now().Add(-time.Hour*24).Format(DATE_FORMAT_ISO), "Start date for searching time entries.")
	until := flag.String("until", time.Now().Format(DATE_FORMAT_ISO), "End date for searching time entries.")
	ansi := flag.Bool("ansi", true, "Whether to format the output with ANSI styles. Used only by the 'text' reporter.")
	flag.Parse()

	err = rules.IsValidISO8601Date(*since, DATE_FORMAT_ISO)
	handleError(err)

	err = rules.IsValidISO8601Date(*until, DATE_FORMAT_ISO)
	handleError(err)

	entries, err := client.Entries(*since, *until)
	handleError(err)

	r, err := getReporter(*output, *ansi, entries)
	handleError(err)

	r.Print()
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getReporter(output string, ansi bool, entries []toggl.TimeEntry) (report.Exportable, error) {
	if output == "text" {
		if ansi {
			return textreporter.NewColorized(entries), nil
		}

		return textreporter.New(entries), nil
	}

	if output == "markdown" {
		return markdownreporter.New(entries), nil
	}

	if output == "csv" {
		return csvreporter.New(entries), nil
	}

	if output == "json" {
		return jsonreporter.New(entries), nil
	}

	return nil, fmt.Errorf("unable to find reporter for output format %s", output)
}
