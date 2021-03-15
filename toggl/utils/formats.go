package utils

import "fmt"

func MillisecondsToHours(seconds int64) float64 {
	return float64(seconds) / 1000 / 60 / 60
}

func FormatDuration(duration float64) string {
	return fmt.Sprintf("%.1f", duration)
}
