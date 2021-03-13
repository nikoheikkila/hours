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
	client := toggl.WithToken(token)

	end := time.Now()
	start := end.Add(-time.Hour * 24 * 7)

	entries, err := client.Entries(start, end)
	handleError(err)

	for _, entry := range entries {
		fmt.Printf("- %.1f hours of project %d in task %s\n", entry.GetDuration(), entry.Pid, entry.Description)
	}
}
