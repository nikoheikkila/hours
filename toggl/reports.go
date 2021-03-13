package toggl

import (
	"time"

	formats "github.com/nikoheikkila/hours/toggl/utils"
)

type TimeEntry struct {
	Id          int       `json:"id"`
	ProjectId   int       `json:"pid"`
	TeamId      int       `json:"tid"`
	UserId      int       `json:"uid"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Updated     time.Time `json:"updated"`
	Duration    int64     `json:"dur"`
	User        string    `json:"user"`
	UseStop     bool      `json:"use_stop"`
	Client      string    `json:"client"`
	Project     string    `json:"project"`
	Color       string    `json:"project_color"`
	HexColor    string    `json:"project_hex_color"`
	Task        string    `json:"task"`
	Billable    float32   `json:"billable"`
	IsBillable  bool      `json:"is_billable"`
	Currency    string    `json:"cur"`
	Tags        []string  `json:"tags"`
}

type DetailedReport struct {
	TotalGrand      int `json:"total_grand"`
	TotalBillable   int `json:"total_billable"`
	TotalCount      int `json:"total_count"`
	PerPage         int `json:"per_page"`
	TotalCurrencies []struct {
		Currency string  `json:"currency"`
		Amount   float32 `json:"amount"`
	} `json:"total_currencies"`
	Data []TimeEntry `json:"data"`
}

// Methods

func (e *TimeEntry) GetDuration() float64 {
	return formats.MillisecondsToHours(e.Duration)
}
