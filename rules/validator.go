package rules

import (
	"fmt"
	"time"
)

func IsValidISO8601Date(date, format string) error {
	_, err := time.Parse(format, date)

	if err != nil {
		return fmt.Errorf("validation error: given date %s is not a valid RFC 3339 string (use the format YYYY-MM-DD)", date)
	}

	return nil
}
