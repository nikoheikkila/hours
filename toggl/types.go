package toggl

import (
	"time"
)

type Entry struct {
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
