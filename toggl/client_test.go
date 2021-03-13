// build +integration

package toggl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	mocks "github.com/nikoheikkila/hours/toggl/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestClientReturnEntriesForTimeRange(t *testing.T) {
	assert := assert.New(t)
	json := `[
		{
			"id": 1,
			"pid": 10,
			"duration": 3600,
			"description": "Working on Hours CLI"
		},
		{
			"id": 2,
			"pid": 10,
			"duration": 1800,
			"description": "Client meeting"
		}
  	]`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       reader,
		}, nil
	}

	toggl := WithToken("some-token")
	entries, err := toggl.Entries(time.Now(), time.Now())

	assert.Nil(err)
	assert.Equal(2, len(entries))

	assert.EqualValues(1, entries[0].Id)
	assert.EqualValues(10, entries[0].Pid)
	assert.EqualValues(3600, entries[0].Duration)
	assert.EqualValues("Working on Hours CLI", entries[0].Description)

	assert.EqualValues(2, entries[1].Id)
	assert.EqualValues(10, entries[1].Pid)
	assert.EqualValues(1800, entries[1].Duration)
	assert.EqualValues("Client meeting", entries[1].Description)
}

func TestClientReturnErrorForMissingToken(t *testing.T) {
	assert := assert.New(t)

	client := WithToken("")
	entries, err := client.Entries(time.Now(), time.Now())

	assert.Nil(entries)
	assert.EqualError(err, "missing Toggl API token for the client")
}
