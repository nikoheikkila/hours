package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidIsoDateReturnsNilOnValidation(t *testing.T) {
	assert := assert.New(t)

	validDates := []string{"2020-02-01", "2021-12-11", "2022-02-28", "2023-07-31"}

	for _, date := range validDates {
		assert.Nil(IsValidISO8601Date(date, "2006-01-02"))
	}
}

func TestInValidIsoDateReturnsErrorOnValidation(t *testing.T) {
	assert := assert.New(t)

	invalidDates := []string{"202-02-00", "2020-02-00", "2021-13-11", "2022-02-30", "2023-07-32"}

	for _, date := range invalidDates {
		v := IsValidISO8601Date(date, "2006-01-02")
		assert.NotNil(v)
		assert.Error(v)
	}
}
