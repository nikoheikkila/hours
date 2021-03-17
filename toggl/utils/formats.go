package utils

import "fmt"

const MILLISECONDS_IN_HOURS float64 = 1000 * 60 * 60

func MillisecondsToHours(seconds int64) float64 {
	return float64(seconds) / MILLISECONDS_IN_HOURS
}

func FormatDuration(duration float64) string {
	return fmt.Sprintf("%.1f", duration)
}
