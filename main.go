package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nikoheikkila/hours/toggl"
)

func env(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Could not read environment variable %s.", key)
	}

	return value
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	token := env("TOGGL_API_TOKEN")
	workspace := env("TOGGL_WORKSPACE_ID")
	client := toggl.New(token, workspace)

	end := time.Now()
	start := time.Now().Add(-time.Hour * 24)

	entries, err := client.Entries(start, end)
	handleError(err)

	for _, entry := range entries {
		fmt.Printf("- %s: %.1f hours of project %s in task %s\n", entry.Start, entry.GetDuration(), entry.Project, entry.Description)
	}
}
