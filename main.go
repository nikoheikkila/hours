package main

import (
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
	"github.com/urfave/cli/v2"
)

const (
	NAME            = "hours"
	VERSION         = "0.3.0"
	AUTHOR_NAME     = "Niko Heikkilä"
	AUTHOR_EMAIL    = "yo@nikoheikkila.fi"
	DESCRIPTION     = "Operate Toggl time entries on the command-line"
	DATE_FORMAT_ISO = "2006-01-02"
)

var (
	configuration *toggl.Configuration
	client        *toggl.TogglClient
)

func init() {
	cli.HelpFlag = &cli.BoolFlag{Name: "help", Aliases: []string{"h"}}
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Aliases: []string{"v", "V"}}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "Version: %s\n", c.App.Version)
	}
}

func main() {
	handleError(app().Run(os.Args))
}

func app() *cli.App {
	return &cli.App{
		Name:     NAME,
		Version:  VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  AUTHOR_NAME,
				Email: AUTHOR_EMAIL,
			},
		},
		Usage:           DESCRIPTION,
		Before:          onBeforeInvocation,
		CommandNotFound: commandNotFound,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "no-ansi",
				Value: false,
				Usage: "Whether to disable the ANSI output in the text reporter.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "list",
				HelpName: NAME + " list",
				Usage: "List time entries in different formats",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "since",
						Aliases:     []string{"s"},
						Value:       time.Now().Add(-time.Hour * 24).Format(DATE_FORMAT_ISO),
						Usage:       "Start date for searching time entries.",
						DefaultText: "yesterday",
					},
					&cli.StringFlag{
						Name:        "until",
						Aliases:     []string{"u"},
						Value:       time.Now().Format(DATE_FORMAT_ISO),
						Usage:       "End date for searching time entries.",
						DefaultText: "today",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format for the reporter.",
					},
				},
				Action: listAction,
			},
		},
	}
}

func onBeforeInvocation(c *cli.Context) error {
	var err error
	configuration, err = toggl.LoadConfiguration()

	if err != nil {
		return err
	}

	client = toggl.New(configuration)
	return nil
}

func listAction(c *cli.Context) error {
	output := c.String("output")
	since := c.String("since")
	until := c.String("until")
	ansiDisabled := c.Bool("no-ansi")

	if err := rules.IsValidISO8601Date(since, DATE_FORMAT_ISO); err != nil {
		return err
	}

	if err := rules.IsValidISO8601Date(until, DATE_FORMAT_ISO); err != nil {
		return err
	}

	entries, err := client.Entries(since, until)
	if err != nil {
		return err
	}

	reporter, err := getReporter(output, ansiDisabled, entries)
	if err != nil {
		return err
	}

	reporter.Print()

	return nil
}

func commandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(c.App.Writer, "Unrecognized command '%s'. Append --help for help.", command)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getReporter(output string, ansiDisabled bool, entries []toggl.TimeEntry) (report.Exportable, error) {
	if output == "" || output == "text" {
		if ansiDisabled {
			return textreporter.New(entries), nil
		}

		return textreporter.NewColorized(entries), nil
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
