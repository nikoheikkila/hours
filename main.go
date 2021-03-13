package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/nikoheikkila/hours/report"
	textreporter "github.com/nikoheikkila/hours/report/text"
	"github.com/nikoheikkila/hours/toggl"
)

func main() {
	token := env("TOGGL_API_TOKEN")
	workspace := env("TOGGL_WORKSPACE_ID")
	client := toggl.New(token, workspace)

	output := flag.String("output", "text", "Output format for the reporter ['text']")
	since := flag.String("since", time.Now().Add(-time.Hour * 24).Format(time.RFC3339), "Start date for searching time entries")
	until := flag.String("until", time.Now().Format(time.RFC3339), "End date for searching time entries")
	flag.Parse()

	entries, err := client.Entries(*since, *until)
	handleError(err)

	r, err := getReporter(*output, entries)
	handleError(err)

	r.Print()
}

func env(key string) string {
	value := os.Getenv(key)

	if value == "" {
		fmt.Printf("Could not read environment variable %s.", key)
		os.Exit(2)
	}

	return value
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getReporter(output string, entries []toggl.TimeEntry) (report.Exportable, error) {
	if output == "text" {
		return textreporter.New(entries), nil
	}

	return nil, fmt.Errorf("unable to find reporter for output format %s", output)
}
