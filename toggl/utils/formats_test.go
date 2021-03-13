package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondsToHours(t *testing.T) {
	tests := []struct {
		seconds int64
		hours   float64
	}{
		{0, 0.0},
		{1800, 0.5},
		{3600, 1.0},
		{7200, 2.0},
		{27000, 7.5},
	}

	for _, test := range tests {
		testcase := fmt.Sprintf("should convert %d seconds to %.1f hours", test.seconds, test.hours)
		t.Run(testcase, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(test.hours, SecondsToHours(test.seconds))
		})
	}
}
