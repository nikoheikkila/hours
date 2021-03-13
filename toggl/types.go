package toggl

import (
	"time"

	formats "github.com/nikoheikkila/hours/toggl/utils"
)

type TimeEntry struct {
    Id int `json:"id"`
    Guid string `json:"guid"`
    Wid int `json:"wid"`
    Pid int `json:"pid"`
    Billable bool `json:"billable"`
    Start time.Time `json:"start"`
    Stop time.Time `json:"stop"`
    Duration int `json:"duration"`
    Description string `json:"description"`
    Duronly bool `json:"duronly"`
    At time.Time `json:"at"`
    Uid int `json:"uid"`
}

func (e *TimeEntry) GetDuration() float64 {
    duration := int64(e.Duration)

    if (duration < 0) {
        return formats.SecondsToHours(time.Now().Unix() + duration)
    }

    return formats.SecondsToHours(duration)
}
